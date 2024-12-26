/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: [
    "../views/**/*.{templ,html}", // Your Go template files
    "../public/**/*.js",   
    'node_modules/preline/dist/*.js',      // JavaScript files in the public folder
  ],
  theme: {
    extend: { },
  },
  plugins: [
    require('preline/plugin')
  ],
};
