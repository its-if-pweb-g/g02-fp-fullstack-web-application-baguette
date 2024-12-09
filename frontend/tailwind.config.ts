import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        primary: "#FFFBE9",
        secondary: {
          light: "#D9CAB3",
          dark: "#AD8B73",
        },
        accent: "#FFC857",
        customRed: "#FF5757",
        customBlack: "#333333",
      },
    },
  },
  plugins: [],
} satisfies Config;
