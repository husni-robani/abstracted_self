<template>
  <AdminLayout>
    <div class="max-w-7xl mx-auto space-y-6">
      <!-- Page header -->
      <div>
        <h1 class="text-3xl font-semibold font-mono mb-1">Visitors</h1>
        <p class="text-gray-500 font-mono">Traffic overview</p>
      </div>

      <!-- Date range controls -->
      <div class="bg-white border border-gray-200 rounded-lg shadow-sm p-4">
        <div
          class="flex flex-col lg:flex-row lg:items-end lg:justify-between gap-4"
        >
          <div class="flex flex-wrap items-end gap-3">
            <div>
              <label
                class="block text-xs uppercase tracking-wide text-gray-500 font-mono mb-1"
              >
                Start
              </label>
              <input
                v-model="rangeStart"
                type="date"
                class="border border-gray-200 rounded-md px-3 py-2 text-sm font-mono text-gray-900 focus:outline-none focus:ring-2 focus:ring-gray-300"
                @input="activePreset = null"
              />
            </div>
            <div>
              <label
                class="block text-xs uppercase tracking-wide text-gray-500 font-mono mb-1"
              >
                End
              </label>
              <input
                v-model="rangeEnd"
                type="date"
                class="border border-gray-200 rounded-md px-3 py-2 text-sm font-mono text-gray-900 focus:outline-none focus:ring-2 focus:ring-gray-300"
                @input="activePreset = null"
              />
            </div>
            <button
              class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono text-sm"
              @click="applyRange"
            >
              Apply
            </button>
          </div>

          <div class="flex flex-wrap gap-2">
            <button
              v-for="preset in presets"
              :key="preset.value"
              class="px-3 py-1.5 rounded-md text-sm font-mono"
              :class="
                activePreset === preset.value
                  ? 'bg-gray-800 text-white'
                  : 'bg-gray-100 text-gray-800 border border-gray-200'
              "
              @click="setPreset(preset.value)"
            >
              {{ preset.label }}
            </button>
          </div>
        </div>
      </div>

      <!-- KPI row -->
      <div class="grid md:grid-cols-4 gap-4">
        <div
          v-for="kpi in kpis"
          :key="kpi.label"
          class="bg-white border border-gray-200 rounded-lg p-4 shadow-sm"
        >
          <div class="text-sm font-mono text-gray-500">{{ kpi.label }}</div>
          <div class="text-2xl font-mono font-bold text-gray-900">
            {{ kpi.value }}
          </div>
          <div v-if="kpi.hint" class="text-xs font-mono text-gray-400 mt-1">
            {{ kpi.hint }}
          </div>
        </div>
      </div>

      <!-- Traffic chart -->
      <div class="bg-white border border-gray-200 rounded-lg shadow-sm">
        <div
          class="px-4 py-3 border-b border-gray-200 flex items-center justify-between gap-3"
        >
          <h2 class="text-lg font-mono font-bold text-gray-900">
            Daily Traffic
          </h2>
          <div class="flex gap-2">
            <button
              v-for="scope in chartScopes"
              :key="scope"
              class="px-3 py-1 rounded-md text-xs font-mono"
              :class="
                chartDays === scope
                  ? 'bg-gray-800 text-white'
                  : 'bg-gray-100 text-gray-800 border border-gray-200'
              "
              @click="setChartDays(scope)"
            >
              {{ scope }}d
            </button>
          </div>
        </div>

        <div class="p-4">
          <!-- Chart loading -->
          <div
            v-if="loadingCounts && chartData.length === 0"
            class="py-16 flex items-center justify-center"
          >
            <p class="text-gray-400 font-mono animate-pulse">
              Loading traffic data…
            </p>
          </div>

          <!-- Chart error -->
          <div
            v-else-if="countsError"
            class="py-12 flex flex-col items-center gap-3 text-center"
          >
            <ExclamationTriangleIcon class="w-8 h-8 text-gray-300" />
            <p class="text-gray-500 font-mono text-sm">
              Couldn't load traffic data
            </p>
            <button
              class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono text-sm"
              @click="retryAll"
            >
              Try again
            </button>
          </div>

          <!-- Bar chart -->
          <template v-else>
            <div class="overflow-x-auto">
              <svg
                :viewBox="`0 0 ${chartWidth} ${CHART.height}`"
                :style="{ minWidth: `${chartWidth}px` }"
                class="w-full h-auto"
                role="img"
                aria-label="Daily visitors bar chart"
              >
                <!-- Gridlines + tick labels -->
                <g v-for="tick in chartTicks" :key="tick.value">
                  <line
                    :x1="CHART.padLeft"
                    :x2="chartWidth - CHART.padRight"
                    :y1="tick.y"
                    :y2="tick.y"
                    stroke="#e5e7eb"
                    stroke-width="1"
                  />
                  <text
                    :x="CHART.padLeft - 8"
                    :y="tick.y + 3"
                    text-anchor="end"
                    fill="#9ca3af"
                    font-size="10"
                    font-family="monospace"
                  >
                    {{ tick.value }}
                  </text>
                </g>

                <!-- Zero baseline -->
                <line
                  :x1="CHART.padLeft"
                  :x2="chartWidth - CHART.padRight"
                  :y1="baselineY"
                  :y2="baselineY"
                  stroke="#9ca3af"
                  stroke-width="1"
                />

                <!-- Bars -->
                <g v-for="bar in chartBars" :key="bar.date" class="group">
                  <rect
                    :x="bar.x"
                    :y="bar.rectY"
                    :width="bar.width"
                    :height="bar.rectHeight"
                    rx="2"
                    :class="
                      bar.stub
                        ? 'fill-gray-300'
                        : 'fill-gray-800 group-hover:fill-gray-700'
                    "
                  >
                    <title>
                      {{ bar.date }}: {{ bar.count }} visitor(s)
                    </title>
                  </rect>
                  <text
                    :x="bar.labelX"
                    :y="bar.rectY - 6"
                    text-anchor="middle"
                    font-size="10"
                    font-family="monospace"
                    :class="
                      bar.stub
                        ? 'fill-gray-400'
                        : 'fill-gray-500 opacity-0 group-hover:opacity-100 transition-opacity'
                    "
                  >
                    {{ bar.count }}
                  </text>
                  <text
                    :x="bar.labelX"
                    :y="baselineY + 16"
                    :transform="`rotate(-40 ${bar.labelX} ${baselineY + 16})`"
                    text-anchor="end"
                    font-size="10"
                    font-family="monospace"
                    :class="
                      bar.index % 2 === 1
                        ? 'fill-gray-500 hidden sm:block'
                        : 'fill-gray-500'
                    "
                  >
                    {{ formatLabel(bar.date) }}
                  </text>
                </g>
              </svg>
            </div>

            <!-- Caption / legend -->
            <p class="text-xs font-mono text-gray-500 mt-2">
              total {{ chartSummary.total }} · peak
              {{ chartSummary.peak }} · {{ chartSummary.zero }} zero-days
            </p>
          </template>
        </div>
      </div>

      <!-- Visitors list -->
      <!-- Loading -->
      <div
        v-if="loadingList && visits.length === 0"
        class="bg-white border border-gray-200 rounded-lg shadow-sm p-4 space-y-3"
      >
        <div class="h-4 w-1/3 bg-gray-200 animate-pulse rounded"></div>
        <div
          v-for="i in 5"
          :key="i"
          class="h-4 bg-gray-200 animate-pulse rounded"
        ></div>
        <p class="text-gray-400 font-mono animate-pulse text-sm">
          Loading traffic data…
        </p>
      </div>

      <!-- Error -->
      <div
        v-else-if="listError"
        class="py-16 flex flex-col items-center gap-4 text-center"
      >
        <ExclamationTriangleIcon class="w-10 h-10 text-gray-300" />
        <p class="text-gray-500 font-mono text-sm">
          Couldn't load traffic data
        </p>
        <button
          class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono text-sm"
          @click="retryAll"
        >
          Try again
        </button>
      </div>

      <!-- Empty -->
      <div
        v-else-if="visits.length === 0"
        class="bg-white border border-dashed border-gray-300 rounded-lg p-12 text-center"
      >
        <UsersIcon class="w-10 h-10 text-gray-300 mx-auto" />
        <p class="mt-3 text-gray-500 font-mono">No visitors in this range</p>
        <button
          class="mt-4 text-gray-800 font-mono text-sm hover:underline"
          @click="resetToLast30"
        >
          Reset to last 30 days
        </button>
      </div>

      <!-- Table -->
      <div
        v-else
        class="bg-white border border-gray-200 rounded-lg shadow-sm overflow-hidden"
      >
        <div
          class="px-4 py-3 border-b border-gray-200 flex items-center justify-between"
        >
          <h2 class="text-lg font-mono font-bold text-gray-900">Visitor Log</h2>
          <span class="text-xs font-mono text-gray-400">
            Maps to data.daily_visits[] + pagination{}
          </span>
        </div>

        <div class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr
                class="text-xs uppercase tracking-wide text-gray-500 border-b border-gray-200"
              >
                <th class="px-4 py-3 font-mono font-medium">Visit Date</th>
                <th class="px-4 py-3 font-mono font-medium">IP Address</th>
                <th class="px-4 py-3 font-mono font-medium">Visitor UUID</th>
                <th class="px-4 py-3 font-mono font-medium hidden md:table-cell">
                  Device (User-Agent)
                </th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="(row, index) in visits"
                :key="row.uuid || index"
                class="border-b border-gray-200 hover:bg-gray-50 text-sm"
              >
                <td class="px-4 py-3 font-mono text-gray-900 whitespace-nowrap">
                  {{ formatDateOnly(row.visit_date) }}
                </td>
                <td class="px-4 py-3 font-mono text-gray-700 whitespace-nowrap">
                  {{ row.ip || "—" }}
                </td>
                <td class="px-4 py-3 font-mono text-gray-600 text-xs break-all">
                  {{ row.uuid || "—" }}
                </td>
                <td
                  class="px-4 py-3 font-mono text-gray-600 text-xs max-w-[24rem] overflow-hidden text-ellipsis whitespace-nowrap hidden md:table-cell"
                  :title="row.device || ''"
                >
                  {{ row.device || "—" }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <div
          class="px-4 py-3 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3"
        >
          <div class="text-xs font-mono text-gray-500">
            Showing {{ rangeStartIndex }}–{{ rangeEndIndex }} of
            {{ pagination.total }}
          </div>

          <div class="flex items-center gap-3">
            <select
              v-model.number="pagination.limit"
              class="border border-gray-200 rounded-md px-2 py-1 text-sm font-mono text-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-300"
              @change="onLimitChange"
            >
              <option :value="10">10</option>
              <option :value="25">25</option>
              <option :value="50">50</option>
            </select>

            <div class="flex items-center gap-1">
              <button
                class="px-3 py-1.5 rounded-md text-sm font-mono border border-gray-200 text-gray-700 hover:bg-gray-50 disabled:opacity-40 disabled:cursor-not-allowed"
                :disabled="pagination.page <= 1"
                @click="goToPage(pagination.page - 1)"
              >
                Prev
              </button>

              <template v-for="(p, index) in visiblePages" :key="index">
                <span
                  v-if="p === '...'"
                  class="px-2 text-gray-400 font-mono text-sm"
                >
                  …
                </span>
                <button
                  v-else
                  class="px-3 py-1.5 rounded-md text-sm font-mono border"
                  :class="
                    p === pagination.page
                      ? 'bg-gray-900 text-white border-gray-900'
                      : 'border-gray-200 text-gray-700 hover:bg-gray-50'
                  "
                  @click="goToPage(p)"
                >
                  {{ p }}
                </button>
              </template>

              <button
                class="px-3 py-1.5 rounded-md text-sm font-mono border border-gray-200 text-gray-700 hover:bg-gray-50 disabled:opacity-40 disabled:cursor-not-allowed"
                :disabled="pagination.page >= pagination.total_pages"
                @click="goToPage(pagination.page + 1)"
              >
                Next
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { computed, onMounted, reactive, ref } from "vue";
import AdminLayout from "../layouts/AdminLayout.vue";
import {
  ExclamationTriangleIcon,
  UsersIcon,
} from "@heroicons/vue/24/outline";

const visits_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_DAILY_VISITS_ENDPOINT;
const counts_endpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_GET_DAILY_VISITS_COUNTS_ENDPOINT;

const token = localStorage.getItem("token");

// ---------------------------------------------------------------------------
// Date helpers (local time, never UTC)
// ---------------------------------------------------------------------------
function formatDate(d) {
  const month = String(d.getMonth() + 1).padStart(2, "0");
  const day = String(d.getDate()).padStart(2, "0");
  return `${d.getFullYear()}-${month}-${day}`;
}

function addDays(date, n) {
  const d = new Date(date);
  d.setDate(d.getDate() + n);
  return d;
}

function parseDateStr(value) {
  const [y, m, d] = String(value).split("-").map(Number);
  return new Date(y, (m || 1) - 1, d || 1);
}

function formatLabel(value) {
  const d = parseDateStr(value);
  return d.toLocaleDateString("en-US", { month: "short", day: "numeric" });
}

function formatDateOnly(value) {
  return value ? String(value).slice(0, 10) : "—";
}

// ---------------------------------------------------------------------------
// Filter state
// ---------------------------------------------------------------------------
const today = new Date();
const rangeStart = ref(formatDate(addDays(today, -6)));
const rangeEnd = ref(formatDate(today));
const activePreset = ref("7");

const presets = [
  { value: "today", label: "Today" },
  { value: "7", label: "Last 7 days" },
  { value: "30", label: "Last 30 days" },
];

const chartScopes = [7, 14, 30];
const chartDays = ref(7);

// ---------------------------------------------------------------------------
// Data state
// ---------------------------------------------------------------------------
const counts = ref([]);
const visits = ref([]);
const loadingCounts = ref(true);
const loadingList = ref(true);
const countsError = ref(false);
const listError = ref(false);

const pagination = reactive({
  page: 1,
  limit: 10,
  total: 0,
  total_pages: 0,
});

// ---------------------------------------------------------------------------
// Fetching
// ---------------------------------------------------------------------------
async function fetchVisits() {
  loadingList.value = true;
  listError.value = false;

  try {
    const query =
      `?start_date=${rangeStart.value}&end_date=${rangeEnd.value}` +
      `&page=${pagination.page}&limit=${pagination.limit}`;

    const res = await fetch(visits_endpoint + query, {
      headers: { Authorization: `Bearer ${token}` },
    });
    const json = await res.json();

    if (!res.ok) {
      throw new Error(json.message || `Request failed: ${res.status}`);
    }

    const data = json.data || {};
    visits.value = Array.isArray(data.daily_visits) ? data.daily_visits : [];

    const meta = data.pagination || {};
    pagination.page = meta.page ?? 1;
    pagination.limit = meta.limit ?? pagination.limit;
    pagination.total = meta.total ?? 0;
    pagination.total_pages = meta.total_pages ?? 0;
  } catch (e) {
    console.error("Failed to fetch visitors:", e);
    listError.value = true;
    visits.value = [];
    pagination.total = 0;
    pagination.total_pages = 0;
  } finally {
    loadingList.value = false;
  }
}

async function fetchCounts() {
  loadingCounts.value = true;
  countsError.value = false;

  try {
    const end = new Date();
    const start = addDays(end, -(chartDays.value - 1));
    const query = `?start_date=${formatDate(start)}&end_date=${formatDate(end)}`;

    const res = await fetch(counts_endpoint + query, {
      headers: { Authorization: `Bearer ${token}` },
    });
    const json = await res.json();

    if (!res.ok) {
      throw new Error(json.message || `Request failed: ${res.status}`);
    }

    counts.value = Array.isArray(json.data) ? json.data : [];
  } catch (e) {
    console.error("Failed to fetch daily visit counts:", e);
    countsError.value = true;
    counts.value = [];
  } finally {
    loadingCounts.value = false;
  }
}

function retryAll() {
  fetchVisits();
  fetchCounts();
}

// ---------------------------------------------------------------------------
// Range controls
// ---------------------------------------------------------------------------
function setPreset(preset) {
  const now = new Date();
  activePreset.value = preset;

  if (preset === "today") {
    rangeStart.value = formatDate(now);
  } else if (preset === "7") {
    rangeStart.value = formatDate(addDays(now, -6));
  } else if (preset === "30") {
    rangeStart.value = formatDate(addDays(now, -29));
  }

  rangeEnd.value = formatDate(now);
  pagination.page = 1;
  fetchVisits();
}

function applyRange() {
  pagination.page = 1;
  fetchVisits();
}

function resetToLast30() {
  setPreset("30");
}

function setChartDays(scope) {
  if (chartDays.value === scope) return;
  chartDays.value = scope;
  fetchCounts();
}

// ---------------------------------------------------------------------------
// Pagination
// ---------------------------------------------------------------------------
function goToPage(page) {
  const max = pagination.total_pages || 1;
  if (page < 1 || page > max || page === pagination.page) return;
  pagination.page = page;
  fetchVisits();
}

function onLimitChange() {
  pagination.page = 1;
  fetchVisits();
}

const rangeStartIndex = computed(() => {
  if (!pagination.total) return 0;
  return (pagination.page - 1) * pagination.limit + 1;
});

const rangeEndIndex = computed(() =>
  Math.min(pagination.page * pagination.limit, pagination.total)
);

const visiblePages = computed(() => {
  const total = pagination.total_pages || 1;
  const current = pagination.page;

  if (total <= 7) {
    return Array.from({ length: total }, (_, i) => i + 1);
  }

  const pages = [1];
  const start = Math.max(2, current - 1);
  const end = Math.min(total - 1, current + 1);

  if (start > 2) pages.push("...");
  for (let i = start; i <= end; i++) pages.push(i);
  if (end < total - 1) pages.push("...");
  pages.push(total);

  return pages;
});

// ---------------------------------------------------------------------------
// Chart geometry
// ---------------------------------------------------------------------------
const CHART = {
  width: 560,
  height: 280,
  padLeft: 44,
  padRight: 16,
  padTop: 24,
  padBottom: 48,
};

const plotHeight = CHART.height - CHART.padTop - CHART.padBottom;
const baselineY = CHART.padTop + plotHeight;

function niceCeil(value) {
  if (value <= 1) return 1;
  const pow = Math.pow(10, Math.floor(Math.log10(value)));
  const norm = value / pow;
  let nice;
  if (norm <= 1) nice = 1;
  else if (norm <= 2) nice = 2;
  else if (norm <= 2.5) nice = 2.5;
  else if (norm <= 5) nice = 5;
  else nice = 10;
  return nice * pow;
}

// Always render exactly `chartDays` bars, ending today (fills zero-count gaps).
const chartData = computed(() => {
  const byDate = new Map();
  (counts.value || []).forEach((row) => {
    if (row && row.visit_date) {
      byDate.set(String(row.visit_date).slice(0, 10), Number(row.count) || 0);
    }
  });

  const now = new Date();
  const out = [];
  for (let i = chartDays.value - 1; i >= 0; i--) {
    const dateStr = formatDate(addDays(now, -i));
    out.push({ date: dateStr, count: byDate.get(dateStr) || 0 });
  }
  return out;
});

const chartWidth = computed(() =>
  Math.max(CHART.width, chartData.value.length * 44)
);

const chartMax = computed(() => {
  const max = chartData.value.reduce((m, d) => Math.max(m, d.count), 0);
  return niceCeil(Math.max(max, 1));
});

const chartTicks = computed(() => {
  const max = chartMax.value || 1;
  const seen = new Set();
  return [max, max / 2, 0]
    .map((v) => Math.round(v))
    .filter((v) => {
      if (seen.has(v)) return false;
      seen.add(v);
      return true;
    })
    .map((v) => ({ value: v, y: baselineY - (v / max) * plotHeight }));
});

const chartBars = computed(() => {
  const data = chartData.value;
  const n = data.length;
  if (n === 0) return [];

  const plotWidth = chartWidth.value - CHART.padLeft - CHART.padRight;
  const slot = plotWidth / n;
  const barWidth = Math.min(slot * 0.55, 28);
  const gap = slot - barWidth;
  const max = chartMax.value || 1;

  return data.map((d, index) => {
    const stub = d.count === 0;
    const rectHeight = stub ? 2 : (d.count / max) * plotHeight;
    const rectY = baselineY - rectHeight;
    const x = CHART.padLeft + index * slot + gap / 2;

    return {
      ...d,
      index,
      x,
      width: barWidth,
      rectY,
      rectHeight,
      stub,
      labelX: x + barWidth / 2,
    };
  });
});

const chartSummary = computed(() => {
  const data = chartData.value;
  return {
    total: data.reduce((acc, d) => acc + d.count, 0),
    peak: data.reduce((m, d) => Math.max(m, d.count), 0),
    zero: data.filter((d) => d.count === 0).length,
  };
});

// ---------------------------------------------------------------------------
// KPIs
// ---------------------------------------------------------------------------
const todayCount = computed(() => {
  const data = chartData.value;
  return data.length ? data[data.length - 1].count : 0;
});

const peak = computed(() =>
  chartData.value.reduce(
    (acc, d) => (d.count > acc.count ? { count: d.count, date: d.date } : acc),
    { count: 0, date: null }
  )
);

const rangeDays = computed(() => {
  if (!rangeStart.value || !rangeEnd.value) return 1;
  const start = parseDateStr(rangeStart.value);
  const end = parseDateStr(rangeEnd.value);
  const diff = Math.round((end - start) / 86400000) + 1;
  return Math.max(1, diff);
});

const avgPerDay = computed(() => {
  if (!pagination.total) return "0.0";
  return (pagination.total / rangeDays.value).toFixed(1);
});

const kpis = computed(() => [
  {
    label: "Total Visitors",
    value: pagination.total,
    hint: `over ${rangeDays.value} day${rangeDays.value === 1 ? "" : "s"}`,
  },
  {
    label: "Today",
    value: chartData.value.length ? todayCount.value : "—",
    hint: "",
  },
  {
    label: "Peak Day",
    value: peak.value.date ? peak.value.count : "—",
    hint: peak.value.date ? formatLabel(peak.value.date) : "",
  },
  { label: "Avg / day", value: avgPerDay.value, hint: "" },
]);

// ---------------------------------------------------------------------------
// Lifecycle
// ---------------------------------------------------------------------------
onMounted(() => {
  fetchVisits();
  fetchCounts();
});
</script>
