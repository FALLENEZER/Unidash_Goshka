/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./internal/template/**/*.{templ,html,go}",
        "./web/**/*.{html,js}"
    ],
    theme: {
        extend: {},
    },
    plugins: [],
}