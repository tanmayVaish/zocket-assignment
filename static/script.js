// This script is only used to add basic interactivity to the HTML templates

// Add a click event listener to the submit button
document
  .querySelector('button[type="submit"]')
  .addEventListener("click", () => {
    // Remove any existing error messages
    const error = document.querySelector(".error-message");
    if (error) {
      error.remove();
    }

    // Get the value of the CSV file path input
    const path = document.querySelector('input[name="path"]').value.trim();

    // If the input is empty, display an error message
    if (!path) {
      const errorMessage = document.createElement("p");
      errorMessage.classList.add("error-message");
      errorMessage.textContent = "Please enter a CSV file path";
      document
        .querySelector("form")
        .insertBefore(
          errorMessage,
          document.querySelector('button[type="submit"]')
        );
    }
  });
