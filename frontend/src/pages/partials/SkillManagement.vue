<template>
  <div class="border rounded-lg p-4 space-y-4">
    <div class="flex items-center justify-between">
      <div class="flex items-center">
        <TagIcon class="w-5 h-5" />
        <span class="px-2 py-1 text-sm rounded">{{ props.type_name }}</span>
      </div>
      <div class="flex space-x-1">
        <button
          @click="showAddSkillModal = true"
          class="flex items-center px-3 py-2 rounded-2xl bg-gray-200 hover:bg-gray-300 space-x-1 hover:cursor-pointer text-gray-900"
        >
          <PlusIcon class="w-4 h-4 color" />
          <span class="text-sm">Add Skill</span>
        </button>
        <button
          @click="removeType"
          class="flex items-center px-3 py-2 rounded-2xl bg-red-100 hover:bg-red-200 space-x-1 hover:cursor-pointer text-red-900"
        >
          <TrashIcon class="w-4 h-4" />
          <span class="text-sm">Delete Type</span>
        </button>
      </div>
    </div>

    <!-- Skills grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
      <div
        v-for="(skill, idx) in props.skill_items"
        :key="idx"
        class="flex items-center justify-between border rounded-lg p-3"
      >
        <div class="flex items-center gap-2">
          <img
            :src="iconEndpoint + '/' + skill.icon_filename"
            alt="icon"
            class="w-6 h-6"
          />
          <span>{{ skill.name }}</span>
        </div>
        <div class="flex items-center gap-3">
          <button
            @click="removeSkill(skill.name)"
            class="text-red-500 hover:text-red-700 hover:cursor-pointer"
          >
            ✕
          </button>
          <!-- Toggle -->
          <label class="inline-flex items-center cursor-pointer">
            <input
              type="checkbox"
              class="sr-only peer"
              :checked="skill.is_most_used"
              @change="() => toggleMostUsed(skill.name)"
            />
            <div
              class="relative w-11 h-6 rounded-full bg-gray-300 transition-colors duration-300 peer-checked:bg-green-400 after:content-[''] after:absolute after:top-0.5 after:left-0.5 after:w-5 after:h-5 after:rounded-full after:bg-white after:shadow after:transition-transform after:duration-300 peer-checked:after:translate-x-5"
            ></div>
          </label>
          <!-- End Toggle -->
        </div>
      </div>
    </div>
    <Modal :show="showAddSkillModal" @close="closeAddSkillModal">
      <div class="space-y-5">
        <h1 class="text-xl font-mono">
          Add skill
          <span class="text-xl font-mono font-semibold">{{
            props.type_name
          }}</span>
        </h1>
        <InputField
          v-model="newSkill.name"
          label="Skill Name"
          placeholder="type new skill ..."
        />
        <InputField
          v-model="newSkill.url"
          label="URL / Documentation"
          placeholder="put the url ..."
        />
        <FileUpload label="Icon File (SVG)" v-model="newSkill.icon_file" />
        <!-- button save -->
        <div class="mt-6 text-right">
          <button
            @click="addSkill"
            class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono hover:cursor-pointer"
          >
            Save
          </button>
        </div>
      </div>
    </Modal>
  </div>
</template>
<script setup>
import { PlusIcon, TagIcon, TrashIcon } from "@heroicons/vue/24/outline";
import { ref } from "vue";
import Modal from "../../components/Modal.vue";
import InputField from "../../components/InputField.vue";
import FileUpload from "../../components/FileUpload.vue";

const iconEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_ICONS_ENDPOINT;
const addSkillEndpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_ADD_PROFILE_SKILL_ENDPOINT;
const removeSkillEndpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_DELETE_PROFILE_SKILL_ENDPOINT;
const toggleMostUsedEndpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_TOGGLE_MOSTUSED_PROFILE_SKILL_ENDPOINT;
const removeSkillTypeEndpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_REMOVE_SKILL_TYPE_ENDPOINT;

const props = defineProps({
  type_name: {
    type: String,
    required: true,
  },
  skill_items: {
    type: Array,
    default: [],
  },
  refresh_skillset: {
    type: Function,
    required: true,
  },
});

const showAddSkillModal = ref(false);

defineEmits();

const newSkill = ref({
  name: "",
  url: "",
  icon_file: null,
});

function closeAddSkillModal() {
  showAddSkillModal.value = false;
}

async function addSkill() {
  const formData = new FormData();

  formData.append("name", newSkill.value.name);
  formData.append("url", newSkill.value.url);
  formData.append("icon_file", newSkill.value.icon_file);
  formData.append("is_most_used", false);
  formData.append("skill_type", props.type_name);

  try {
    const token = localStorage.getItem("token");
    const res = await fetch(addSkillEndpoint, {
      method: "POST",
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(`Create project failed: ${resJson.message}`);
    }

    showAddSkillModal.value = false;
    await props.refresh_skillset();
  } catch (e) {
    console.error(e);
  }
}

async function removeSkill(skill_name) {
  if (!confirm("Are you sure want to delete skill ?")) {
    return;
  }

  try {
    const token = localStorage.getItem("token");
    const formData = new FormData();

    formData.append("skill_name", skill_name);

    const res = await fetch(removeSkillEndpoint, {
      method: "DELETE",
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error("failed to remove skill");
    }

    await props.refresh_skillset();
  } catch (e) {
    console.error(e);
  }
}

async function toggleMostUsed(skill_name) {
  try {
    const token = localStorage.getItem("token");
    const formData = new FormData();
    formData.append("skill_name", skill_name);
    formData.append("type_name", props.type_name);
    const res = await fetch(toggleMostUsedEndpoint, {
      method: "POST",
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (!res.ok) {
      throw new Error("Failed to toggle skill most used state");
    }
    await props.refresh_skillset();
  } catch (e) {
    console.error(e);
  }
}

async function removeType() {
  if (!confirm("Are you sure want to delete type including it's skills ?")) {
    return;
  }

  try {
    const token = localStorage.getItem("token");
    const formData = new FormData();

    formData.append("type_name", props.type_name);

    const res = await fetch(removeSkillTypeEndpoint, {
      method: "DELETE",
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error("Failed to remove skill type");
    }

    await props.refresh_skillset();
  } catch (e) {
    console.error(e);
  }
}
</script>
