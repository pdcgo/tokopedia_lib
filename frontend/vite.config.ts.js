// vite.config.ts
import react from "@vitejs/plugin-react";
import { customAlphabet } from "nanoid";
import { defineConfig } from "vite";
var vite_config_default = defineConfig({
  base: "",
  plugins: [react({
    jsxImportSource: "@emotion/react"
  })],
  build: {
    rollupOptions: {
      output: {
        chunkFileNames() {
          const nanoid = customAlphabet("7AopXkhgrtyeQ", 9);
          const name = nanoid();
          return `assets/js/nox_${name}.js`;
        },
        entryFileNames() {
          const nanoid = customAlphabet("7AopXkhgrtyeQ", 9);
          const name = nanoid();
          return `assets/js/nox_${name}.js`;
        },
        assetFileNames: (assetInfo) => {
          const ext = assetInfo.name.split(".")[assetInfo.name.split(".").length - 1];
          let extType = assetInfo.name.split(".")[assetInfo.name.split(".").length - 1];
          if (/png|jpe?g|svg|gif|tiff|bmp|ico/i.test(extType)) {
            extType = "images";
          }
          const nanoid = customAlphabet("7AopXkhgrtyeQ", 9);
          const name = nanoid();
          return `assets/${extType}/nox_${name}.${ext}`;
        }
      }
    }
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCByZWFjdCBmcm9tICdAdml0ZWpzL3BsdWdpbi1yZWFjdCdcclxuaW1wb3J0IHsgY3VzdG9tQWxwaGFiZXQgfSBmcm9tIFwibmFub2lkXCJcclxuaW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSAndml0ZSdcclxuXHJcbi8vIGh0dHBzOi8vdml0ZWpzLmRldi9jb25maWcvXHJcbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZyh7XHJcbiAgYmFzZTogJycsXHJcbiAgcGx1Z2luczogW3JlYWN0KHtcclxuICAgIGpzeEltcG9ydFNvdXJjZTogXCJAZW1vdGlvbi9yZWFjdFwiXHJcbiAgfSldLFxyXG4gIGJ1aWxkOiB7XHJcbiAgICByb2xsdXBPcHRpb25zOiB7XHJcbiAgICAgIG91dHB1dDoge1xyXG4gICAgICAgIGNodW5rRmlsZU5hbWVzKCkge1xyXG4gICAgICAgICAgY29uc3QgbmFub2lkID0gY3VzdG9tQWxwaGFiZXQoXCI3QW9wWGtoZ3J0eWVRXCIsIDkpXHJcbiAgICAgICAgICBjb25zdCBuYW1lID0gbmFub2lkKClcclxuICAgICAgICAgIHJldHVybiBgYXNzZXRzL2pzL25veF8ke25hbWV9LmpzYFxyXG4gICAgICAgIH0sXHJcbiAgICAgICAgZW50cnlGaWxlTmFtZXMoKSB7XHJcbiAgICAgICAgICBjb25zdCBuYW5vaWQgPSBjdXN0b21BbHBoYWJldChcIjdBb3BYa2hncnR5ZVFcIiwgOSlcclxuICAgICAgICAgIGNvbnN0IG5hbWUgPSBuYW5vaWQoKVxyXG4gICAgICAgICAgcmV0dXJuIGBhc3NldHMvanMvbm94XyR7bmFtZX0uanNgXHJcbiAgICAgICAgfSxcclxuICAgICAgICBhc3NldEZpbGVOYW1lczogKGFzc2V0SW5mbykgPT4ge1xyXG4gICAgICAgICAgY29uc3QgZXh0ID0gYXNzZXRJbmZvLm5hbWUuc3BsaXQoJy4nKVthc3NldEluZm8ubmFtZS5zcGxpdCgnLicpLmxlbmd0aCAtIDFdO1xyXG5cclxuICAgICAgICAgIGxldCBleHRUeXBlID0gYXNzZXRJbmZvLm5hbWUuc3BsaXQoJy4nKVthc3NldEluZm8ubmFtZS5zcGxpdCgnLicpLmxlbmd0aCAtIDFdO1xyXG4gICAgICAgICAgaWYgKC9wbmd8anBlP2d8c3ZnfGdpZnx0aWZmfGJtcHxpY28vaS50ZXN0KGV4dFR5cGUpKSB7XHJcbiAgICAgICAgICAgIGV4dFR5cGUgPSAnaW1hZ2VzJztcclxuICAgICAgICAgIH1cclxuXHJcbiAgICAgICAgICBjb25zdCBuYW5vaWQgPSBjdXN0b21BbHBoYWJldChcIjdBb3BYa2hncnR5ZVFcIiwgOSlcclxuICAgICAgICAgIGNvbnN0IG5hbWUgPSBuYW5vaWQoKVxyXG5cclxuICAgICAgICAgIHJldHVybiBgYXNzZXRzLyR7ZXh0VHlwZX0vbm94XyR7bmFtZX0uJHtleHR9YDtcclxuICAgICAgICB9LFxyXG4gICAgICB9XHJcbiAgICB9XHJcbiAgfVxyXG59KVxyXG4iXSwKICAibWFwcGluZ3MiOiAiO0FBQUE7QUFDQTtBQUNBO0FBR0EsSUFBTyxzQkFBUSxhQUFhO0FBQUEsRUFDMUIsTUFBTTtBQUFBLEVBQ04sU0FBUyxDQUFDLE1BQU07QUFBQSxJQUNkLGlCQUFpQjtBQUFBLEVBQ25CLENBQUMsQ0FBQztBQUFBLEVBQ0YsT0FBTztBQUFBLElBQ0wsZUFBZTtBQUFBLE1BQ2IsUUFBUTtBQUFBLFFBQ04saUJBQWlCO0FBQ2YsZ0JBQU0sU0FBUyxlQUFlLGlCQUFpQixDQUFDO0FBQ2hELGdCQUFNLE9BQU8sT0FBTztBQUNwQixpQkFBTyxpQkFBaUI7QUFBQSxRQUMxQjtBQUFBLFFBQ0EsaUJBQWlCO0FBQ2YsZ0JBQU0sU0FBUyxlQUFlLGlCQUFpQixDQUFDO0FBQ2hELGdCQUFNLE9BQU8sT0FBTztBQUNwQixpQkFBTyxpQkFBaUI7QUFBQSxRQUMxQjtBQUFBLFFBQ0EsZ0JBQWdCLENBQUMsY0FBYztBQUM3QixnQkFBTSxNQUFNLFVBQVUsS0FBSyxNQUFNLEdBQUcsRUFBRSxVQUFVLEtBQUssTUFBTSxHQUFHLEVBQUUsU0FBUztBQUV6RSxjQUFJLFVBQVUsVUFBVSxLQUFLLE1BQU0sR0FBRyxFQUFFLFVBQVUsS0FBSyxNQUFNLEdBQUcsRUFBRSxTQUFTO0FBQzNFLGNBQUksa0NBQWtDLEtBQUssT0FBTyxHQUFHO0FBQ25ELHNCQUFVO0FBQUEsVUFDWjtBQUVBLGdCQUFNLFNBQVMsZUFBZSxpQkFBaUIsQ0FBQztBQUNoRCxnQkFBTSxPQUFPLE9BQU87QUFFcEIsaUJBQU8sVUFBVSxlQUFlLFFBQVE7QUFBQSxRQUMxQztBQUFBLE1BQ0Y7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUNGLENBQUM7IiwKICAibmFtZXMiOiBbXQp9Cg==
