/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    "./node_modules/flowbite/**/*.js"
  ],
  theme: {
    extend: {
      colors: {
        //Primary Color
        BlackColor :'#0F0402',
        AccentColor: '#35665B',
        SecondColor: '#5A8D80',
        SupportColor1: '#81B5A8',
        SupportColor2: '#AADFD1',
        DomlantColor: '#FEFEFE',
        //Neutral Color
        IntroColor: '#5BA4FC',
        DangerColor: '#FC3D39',
        WarningColor: '#FC9C30',
        SuccessColor: '#049F67',
        DisableColor: '#F8F8F8',

      },
      typography: {
        DEFAULT: {
          '.Heading1': {
            fontWeight: '600',
            fontSize: '6rem',
            lineHeight: '1',
          },
          '.Heading2': {
            fontWeight: '300',
            fontSize: '3.75rem',
            lineHeight: 'normal',
          },
          '.Heading3': {
            fontWeight: '600',
            fontSize: '3rem',
            lineHeight: 'normal',
          },
          '.Heading4': {
            fontWeight: '400',
            fontSize: '2.125rem',
            lineHeight: 'normal',
          },
          '.Heading5': {
            fontWeight: '400',
            fontSize: '1.5rem',
            lineHeight: 'normal',
          },
          '.Heading6': {
            fontWeight: 'bold',
            fontSize: '1.25rem',
            lineHeight: 'normal',
          },
          '.Subtitle1': {
            fontWeight: '500',
            fontSize: '1.25rem',
            lineHeight: 'normal',
          },
          '.Subtitle2': {
            fontWeight: '500',
            fontSize: '1rem',
            lineHeight: 'normal',
          },
          '.Body1': {
            fontWeight: '400',
            fontSize: '1rem',
            lineHeight: 'normal',
          },
          '.Body2': {
            fontWeight: '400',
            fontSize: '0.875rem',
            lineHeight: 'normal',
          },
        }
      }
    },
  },
  plugins: [
    require('@tailwindcss/typography'), // Kích hoạt plugin Typography
  ],
}

