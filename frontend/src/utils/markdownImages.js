// Shared helpers for Blog CMS image URL handling (AD-5 / AD-11).
//
// The backend stores image URLs relative to the API root (`/images/{id}`). On the
// client we prefix them with VITE_API_URL so the browser can resolve them, but we
// never rewrite them in the stored markdown (the backend keeps the relative form).

const VITE_API_URL = import.meta.env.VITE_API_URL || "";

/**
 * Rewrites relative `/images/` URLs to absolute API-backed URLs in an HTML fragment.
 * Used as the md-editor-v3 `sanitize` callback so that live previews (and the public
 * reading page's MdPreview) resolve embedded images against the API origin.
 *
 * Idempotent by construction: after the first pass the match sequence `src="` /
 * `href="` is followed by `<VITE_API_URL>/images/`, which no longer matches the
 * `(src|href)="\/images/` pattern.
 *
 * @param {string} html
 * @returns {string}
 */
export function resolveImageUrls(html) {
  if (!html || typeof html !== "string") return html;

  const apiImageBase = VITE_API_URL + "/images/";

  return html.replace(/(src|href)="\/images\//g, `$1="${apiImageBase}`);
}

/**
 * Extracts a unique, ordered list of image ids referenced in markdown content as
 * `/images/{id}`. Used to send `content_image_ids` on create and as a
 * full-replacement list on update (AD-11).
 *
 * @param {string} markdown
 * @returns {number[]}
 */
export function extractImageIds(markdown) {
  if (!markdown || typeof markdown !== "string") return [];

  const ids = [];
  const regex = /\/images\/(\d+)/g;
  let match = null;

  while ((match = regex.exec(markdown)) !== null) {
    const id = Number(match[1]);
    if (!ids.includes(id)) ids.push(id);
  }

  return ids;
}
