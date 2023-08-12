class TribeEditorElement extends HTMLElement {
  elements = [];
  constructor() {
    super();

    this.id = "tribe-editor-elm";
    this.className = "tribe-editor-elm";
    window.__editor = this;
  }

  connectedCallback() {
    this.attachNewEditor();
  }

  attachNewEditor() {
    const editorElm = document.createElement("tribe-item");
    this.appendChild(editorElm);
    this.elements.push(editorElm);

    editorElm.addEventListener("press-enter", (e) => {
      console.log("press-enter", e);
    });
  }
}

export default TribeEditorElement;
