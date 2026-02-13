export function addUploadButtonListener() {
    console.log("WORKING");
    document.addEventListener('DOMContentLoaded', function() {
        const pubBtn = document.getElementById('uploadButton');
        pubBtn.addEventListener('click', uploadHandle);
    });
}

async function uploadHandle() {
    let doc = document.getElementById('fileToUpload').files[0];

    const formData = new FormData();

    formData.append('media', doc);
    formData.append('description', 'My PDF descrption');
    formData.append('title', 'My Title');

    try {
        const response = await fetch('http://localhost:8080/upload', {
            method: 'POST',
            body: formData
        });

        const data = await response.arrayBuffer();

        if (data.error) {
            alert('Error: ' + data.error);
        } else {
            alert('uploaded')
        }
    } catch (error) {
        alert('Failed to connect to server: ' + error.message)
    }
}
