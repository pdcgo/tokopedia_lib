import react from '@vitejs/plugin-react'
import { customAlphabet } from "nanoid"
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  base: '',
  plugins: [react()],
  build: {
    rollupOptions: {
      output: {
        chunkFileNames() {
          const nanoid = customAlphabet("7AopXkhgrtyeQ", 9)
          const name = nanoid()
          return `assets/js/nox_${name}.js`
        },
        entryFileNames() {
          const nanoid = customAlphabet("7AopXkhgrtyeQ", 9)
          const name = nanoid()
          return `assets/js/nox_${name}.js`
        },
        assetFileNames: (assetInfo) => {
          const ext = assetInfo.name.split('.')[assetInfo.name.split('.').length - 1];

          let extType = assetInfo.name.split('.')[assetInfo.name.split('.').length - 1];
          if (/png|jpe?g|svg|gif|tiff|bmp|ico/i.test(extType)) {
            extType = 'images';
          }

          const nanoid = customAlphabet("7AopXkhgrtyeQ", 9)
          const name = nanoid()

          return `assets/${extType}/nox_${name}.${ext}`;
        },
      }
    }
  }
})
