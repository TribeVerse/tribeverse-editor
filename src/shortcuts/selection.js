import Mousetrap from "mousetrap";

Mousetrap.bind("command+down", function (e) {
  e.preventDefault(); // Prevent the default browser action
  console.log("Ctrl + N pressed"); // Replace with your desired action

  if (__editor) {
    __editor.attachNewEditor();
  }
});
