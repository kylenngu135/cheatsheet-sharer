const API_URL = 'http://localhost:8080';

async function loadFiles() {
    const response = await fetch(API_URL + '/api/files');
    const files = await response.json();
    displayFiles(files);
}

function displayFiles(files) {
    const list = document.getElementById('file-list');

    if (files) {
        files.forEach(file => {
            const li = document.createElement('li');
            li.textContent = file.name;
            li.onclick = function() {
              viewFile(file);
            };
            list.append(li);
        });
    }
}

function viewFile(file) {
    const modal = document.getElementById('modal');
    const viewer = document.getElementById('viewer');

    viewer.innerHTML = '';

    if (file.type === 'image') {
        const img = document.createElement('img');
        img.src = API_URL + '/api/files/' + file.name;
        viewer.appendChild(img);
    } else if (file.type === 'pdf') {

        const iframe = document.createElement('iframe');
        iframe.src = API_URL + '/api/files/' + file.name;
        iframe.height = '800';
        iframe.width = '600';
        viewer.appendChild(iframe);
    }

    modal.style.display = 'block';
}

function closeModal() {
    const modal = document.getElementById('modal');
    modal.style.display = 'none';
}

loadFiles();
