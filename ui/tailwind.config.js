/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
    "./public/**/*.html",
    "./node_modules/flowbite-react/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui"),  require('@tailwindcss/typography'),],
  // daisyUI config (optional)
  daisyui: {
    styled: true,
    themes: ["light", "dark", "night"],
    base: true,
    utils: true,
    logs: true,
    rtl: false,
    prefix: "",
  },
}
