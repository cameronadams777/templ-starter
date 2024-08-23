/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['views/**/*.templ'],
  theme: {
    extend: {
      keyframes: {
        slideInFromTop: {
          '0%': { transform: 'translateY(-100%)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' },
        },
        slideOutToTop: {
          '0%': { transform: 'translateY(0)', opacity: '1' },
          '100%': { transform: 'translateY(-100%)', opacity: '0' },
        },
      },
      animation: {
        slideInFromTop: 'slideInFromTop 0.5s ease-out forwards',
        slideOutToTop: 'slideOutToTop 0.5s ease-in forwards',
      },
    },
  },
  plugins: [],
}

