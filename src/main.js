import "./shortcuts/selection";

import TribeItemElement from "./elements/item";
import TribeEditorElement from "./elements/editor";

if (!customElements.get("tribe-item")) {
  customElements.define("tribe-item", TribeItemElement);
}

if (!customElements.get("tribe-editor")) {
  customElements.define("tribe-editor", TribeEditorElement);
}
