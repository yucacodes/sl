function submitHandle(event) {
  const submitButton = document.getElementById('submitButton');
  const inputLink = document.getElementById('inputLink');

  if (submitButton && inputLink) {
    submitButton.disabled = true;
    inputLink.disabled = true;
    submitButton.classList.add("loading");
    console.log("Entro al boton...")
    setTimeout(() => {
      // Restore the button state  
      //after the operation is done 
      submitButton.disabled = false;
      inputLink.disabled = false;
      submitButton.classList.remove("loading");
    }, 4000);
  }

  console.log(`Se ha generado el come monda enlace corto! Timestamp: ${event.timeStamp}`);
  event.preventDefault();
}


window.addEventListener("DOMContentLoaded", (event) => {

  const urlForm = document.getElementById('form');


  if (urlForm) {
    urlForm.addEventListener('submit', submitHandle, false);
    console.log("Si es valido...")
    urlForm.reset();
  }
})