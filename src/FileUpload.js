async function upload_handle() {
    let doc = document.getElementById('fileToUpload').value;

    console.log(doc);
}

document.addEventListener('DOMContentLoaded', function() {
    const pubBtn = document.getElementById('uploadButton');
    pubBtn.addEventListener('click', upload_handle);
});


