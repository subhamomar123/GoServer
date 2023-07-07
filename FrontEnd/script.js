const form = document.getElementById('commandForm');
const outputContainer = document.getElementById('outputContainer');
const outputText = document.getElementById('outputText');

form.addEventListener('submit', (event) => {
    event.preventDefault();

    const command = document.getElementById('commandInput').value;

    fetch('http://localhost:8080/api/cmd?command=' + encodeURIComponent(command))
        .then(response => response.text())
        .then(data => {
            outputText.textContent = data;
            outputContainer.style.display = 'block';
        })
        .catch(error => {
            outputText.textContent = 'Error: ' + error.message;
            outputContainer.style.display = 'block';
        });
});
