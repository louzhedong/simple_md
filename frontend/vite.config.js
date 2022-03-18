import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path';

const resolveDir = dir => path.resolve(__dirname, dir);

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  plugins: [vue()],
  resolve: {
    alias: {
      '@src': resolveDir('./src'),
      '@assets': resolveDir('./src/assets'),
      '@common': resolveDir('./src/common'),
      '@components': resolveDir('./src/components')
    }
  }
})
