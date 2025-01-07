/** @type {import('tailwindcss').Config} */
module.exports = {
  purge: false,
  darkMode: 'class',
  content: [
    "../views/**/*.{templ,html}", // Your Go template files
    "../public/**/*.js",
    "node_modules/preline/dist/*.js",      // JavaScript files in the public folder
    "node_modules/flyonui/dist/js/*.js"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('preline/plugin'),
    require("flyonui"),
    require("flyonui/plugin")
  ],
  flyonui: {
    themes: "light"
  }
};
