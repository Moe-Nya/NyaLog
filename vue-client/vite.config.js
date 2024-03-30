import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import AutoImports from "unplugin-auto-import/vite";
import { dirResolver, DirResolverHelper } from "vite-auto-import-resolvers";
import {
  NaiveUiResolver,
} from 'unplugin-vue-components/resolvers'

// your plugin installation
Components({
  resolvers: [
    NaiveUiResolver(),
  ],
})

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), DirResolverHelper(), AutoImports({imports : ['vue', 'vue-router', 'pinia'], resolvers: [dirResolver(), NaiveUiResolver(),]}), Components()],
})