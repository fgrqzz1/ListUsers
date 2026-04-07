async function loadUsers() {
    const statusDiv = document.getElementById('status');
    const table = document.getElementById('userTable');
    const tbody = document.getElementById('userBody');

    try {
        statusDiv.innerText = "Загрузка...";
        
        const response = await fetch('http://localhost:8080/users');

        // Проверяем, что сервер ответил успешно (код 200-299)
        if (!response.ok) {
            throw new Error(`Ошибка сервера: ${response.status}`);
        }

        const users = await response.json();

        tbody.innerHTML = '';
        users.forEach(user => {
            const row = `<tr>
                <td>${user.id}</td>
                <td>${user.name}</td>
                <td>${user.email}</td>
            </tr>`;
            tbody.innerHTML += row;
        });

        table.style.display = 'table';
        statusDiv.innerText = "";

    } catch (error) {
        console.error("Fetch error:", error);
        statusDiv.innerHTML = `<p class="error">Сервис временно недоступен. Попробуйте позже.</p>`;
        table.style.display = 'none';
    }
}

loadUsers();