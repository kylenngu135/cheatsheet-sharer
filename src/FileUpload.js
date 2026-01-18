async function upload_handle() {
    let doc = document.getElementById('fileToUpload').value;

    try {
        const response = await fetch('http://localhost:8080/upload', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                document: doc,
            })
        });

        const data = await response.json();

        if (data.error) {
            alert('Error: ' + data.error);
        } else {
            alert('SUCCESS!')
        }
    } catch (error) {
        alert('Failed to connect to server: ' + error.message)
    }
}

document.addEventListener('DOMContentLoaded', function() {
    const pubBtn = document.getElementById('uploadButton');
    pubBtn.addEventListener('click', upload_handle);
});
