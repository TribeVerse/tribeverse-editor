import commonjs from "@rollup/plugin-commonjs";
import resolve from "@rollup/plugin-node-resolve";
import peerDepsExternal from "rollup-plugin-peer-deps-external";
import terser from "@rollup/plugin-terser";

export default {
  input: "src/main.js",
  output: [
    {
      file: "dist/tribe-editor.umd.js",
      format: "umd",
      name: "TribeEditor",
    },
    {
      file: "dist/tribe-editor.esm.js",
      format: "esm",
    },
  ],
  plugins: [peerDepsExternal(), resolve(), commonjs(), terser()],
};
