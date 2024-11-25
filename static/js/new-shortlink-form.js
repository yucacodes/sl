
function initNewShortLinkForm(form) {
  const submitButton = form.querySelector('button[type="submit"]')
  const inputs = form.querySelectorAll('input')
  form.addEventListener('submit', (evt) => {
    evt.preventDefault()//ojo
    for (const input of inputs) {
      input.readOnly = true
    }
    submitButton.disabled = true
    submitButton.classList.add("loading");
  })
}

window.addEventListener("DOMContentLoaded", (event) => {
  const form = document.getElementById('new-shortlink-form')

  if (form) {
    initNewShortLinkForm(form)
  }
})