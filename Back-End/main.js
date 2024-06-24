// Espera que todo o conteúdo do DOM seja carregado antes de executar o código
document.addEventListener('DOMContentLoaded', () => {
    // Seleciona o campo de entrada de busca e o elemento principal do conteúdo
    const searchInput = document.querySelector('.search-input');
    const mainContent = document.querySelector('main');

    // Chama a função para buscar promoções ao carregar a página
    fetchPromotions();

    // Adiciona um ouvinte de evento para o campo de busca para detectar quando a tecla 'Enter' é pressionada
    searchInput.addEventListener('keypress', function (e) {
        if (e.key === 'Enter') {
            const searchQuery = searchInput.value; // Obtém o valor da busca
            searchPromotions(searchQuery); // Chama a função de busca de promoções com o valor da busca
        }
    });

    // Função para buscar promoções do servidor
    function fetchPromotions() {
        fetch('http://localhost:8081/promotions/report') // Faz uma requisição GET para o endpoint de promoções
            .then(response => response.json()) // Converte a resposta em JSON
            .then(data => displayPromotions(data)) // Chama a função para exibir as promoções
            .catch(error => console.error('Error fetching promotions:', error)); // Lida com erros na requisição
    }

    // Função para exibir as promoções na página
    function displayPromotions(promotions) {
        mainContent.innerHTML = '<h3>Promoções</h3>'; // Adiciona um título ao conteúdo principal
        promotions.forEach(promo => { // Para cada promoção, cria um elemento de div e adiciona ao conteúdo principal
            const promoElement = document.createElement('div');
            promoElement.innerHTML = `
                <h4>${promo.ID} - ${promo.Codigo}</h4>
                <p>Desconto: ${promo.Desconto}%</p>
                <p>Produtos: ${promo.Produtos.join(', ')}</p>
                <p>Validade: ${new Date(promo.ValidadeInicio).toLocaleDateString()} - ${new Date(promo.ValidadeFim).toLocaleDateString()}</p>
            `;
            mainContent.appendChild(promoElement); // Adiciona o elemento da promoção ao conteúdo principal
        });
    }

    // Função para buscar promoções com base em uma consulta de busca
    function searchPromotions(query) {
        fetchPromotions(); // No momento, apenas chama a função fetchPromotions; pode ser expandido para buscar com base na consulta
    }

    // Função para aplicar uma promoção
    function applyPromotion(codigo, produtos, totalCompra) {
        fetch('http://localhost:8081/promotions/apply', {
            method: 'POST', // Define o método HTTP como POST
            headers: {
                'Content-Type': 'application/json' // Define o cabeçalho da requisição como JSON
            },
            body: JSON.stringify({ codigo, produtos, totalCompra }) // Converte o corpo da requisição para JSON
        })
            .then(response => response.json()) // Converte a resposta em JSON
            .then(data => {
                console.log('Promotion applied:', data); // Loga a resposta no console
                alert(`Desconto: R$${data.desconto}, Total: R$${data.total}`); // Mostra um alerta com os detalhes da promoção aplicada
            })
            .catch(error => console.error('Error applying promotion:', error)); // Lida com erros na requisição
    }

    // Chama a função applyPromotion com valores de teste
    applyPromotion('PROMO10', ['produto1', 'produto2'], 100.00);
});
