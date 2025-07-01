async function generateAscii() {
    const fileInput = document.getElementById('image-upload');
    const xInput = document.getElementById('x');
    const yInput = document.getElementById('y');
    const output = document.getElementById('output');

    if (!fileInput.files[0]) {
        output.value = "Seleziona un file immagine.";
        return;
    }

    const formData = new FormData();
    formData.append('file', fileInput.files[0]);

    const x = encodeURIComponent(xInput.value);
    const y = encodeURIComponent(yInput.value);

    try {
        const response = await fetch(`http://127.0.0.1:8080/.netlify/functions/generate?x=${x}&y=${y}`, {
            method: 'POST',
            body: formData
        });

        if (!response.ok) {
            output.value = "Errore dal server: " + response.statusText;
            return;
        }

        const text = await response.text();
        output.value = text;
    } catch (err) {
        output.value = "Errore di rete: " + err;
    }
}