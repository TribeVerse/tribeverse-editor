import "trix";
import debounce from "debounce";

class TribeItemElement extends HTMLElement {
  editorInstance = null;
  constructor() {
    super();
  }

  connectedCallback() {
    this.className = "tribe-item-elm";
    const editorElm = document.createElement("trix-editor");
    editorElm.className = editorElm.className + " tribe-item";

    this.editorInstance = editorElm;
    this.appendChild(editorElm);
    this.focus();
    this.registerEvents();
  }

  registerEvents() {
    this.editorInstance.addEventListener("trix-change", () => {
      console.log("change");
    });

    this.editorInstance.addEventListener("keypress", (e) => {
      if (e.key === "Enter") {
        this.dispatchEnterEvent();
      }
    });

    this.editorInstance.addEventListener("trix-focus", (e) => {
      //   if (this.editorInstance && this.editorInstance.toolbarElement) {
      //     this.editorInstance.toolbarElement.style.display = "block";
      //   }
    });

    this.editorInstance.addEventListener(
      "trix-blur",
      debounce((e) => {
        // if (this.editorInstance && this.editorInstance.toolbarElement) {
        //   this.editorInstance.toolbarElement.style.display = "none";
        // }
      }, 1000)
    );
  }

  dispatchEnterEvent() {
    const event = new CustomEvent("press-enter", {
      detail: {
        content: this.editorInstance.value,
      },
    });
    this.dispatchEvent(event);
  }
}

export default TribeItemElement;
