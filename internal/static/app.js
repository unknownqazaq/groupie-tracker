// Функция для выполнения fetch запроса и отображения информации об артисте
function fetchArtistInfo() {
    fetch('/artists')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            // Вывод данных в консоль для проверки
            console.log(data);

            // Получение элемента, в который будем добавлять информацию об артистах
            const artistInfoDiv = document.getElementById('artistInfo');

            // Цикл для обхода всех артистов в полученных данных
            data.forEach(artist => {
                // Создаем новый div для каждого артиста
                const artistDiv = document.createElement('div');
                artistDiv.classList.add('col'); // Добавляем класс для столбца Bootstrap

                // Заполняем div информацией об артисте
                artistDiv.innerHTML = `
            <div class="artist-info text-center">
                <h2>${artist.name}</h2>
              <button onclick="infoArtist(${artist.id})" class="custom-button"> <img src="${artist.image}" alt="${artist.name}"> </button>

                <p class="d-none">${artist.id}</p>
            </div>
        `;

                // Добавляем созданный div с информацией об артисте в общий контейнер
                artistInfoDiv.appendChild(artistDiv);
            });
        })
        .catch(error => {
            console.error('There was a problem with your fetch operation:', error);
        });

}

function infoArtist(artistId) {
    // Перенаправляем пользователя на страницу с информацией об артисте
    window.location.href = `/artistInfo?id=${artistId}`;
}


// Вызов функции для получения информации об артисте при загрузке страницы
window.onload = fetchArtistInfo;