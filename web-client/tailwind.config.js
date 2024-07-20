import { nextui } from '@nextui-org/theme';

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
    './node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['var(--font-opensans)'],
        mono: ['var(--font-geist-mono)'],
      },
      colors: {
        mainLight: '#60BEEB',
        mainDark: '#115E83',
      },
    },
  },
  darkMode: 'class',
  plugins: [nextui()],
};
