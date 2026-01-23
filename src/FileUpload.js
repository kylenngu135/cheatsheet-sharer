async function upload_handle() {
    let doc = document.getElementById('fileToUpload').files[0];

    const formData = new FormData();

    formData.append('file', doc);
    formData.append('descript', 'My PDF descrption');
    formData.append('title', 'My Title');

    try {
        const response = await fetch('http://localhost:8080/upload', {
            method: 'POST',
            body: formData
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
