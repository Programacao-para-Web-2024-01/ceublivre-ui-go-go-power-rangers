
document.addEventListener('DOMContentLoaded', () => {
    const searchInput = document.querySelector('.search-input');
    const mainContent = document.querySelector('main');

    fetchPromotions();

    searchInput.addEventListener('keypress', function (e) {
        if (e.key === 'Enter') {
            const searchQuery = searchInput.value; 
            searchPromotions(searchQuery); 
        }
    });

    function fetchPromotions() {
        fetch('http://localhost:8081/promotions/report') 
            .then(response => response.json()) 
            .then(data => displayPromotions(data)) 
            .catch(error => console.error('Error fetching promotions:', error)); 
    }

    function displayPromotions(promotions) {
        mainContent.innerHTML = '<h3>Promoções</h3>'; 
        promotions.forEach(promo => { 
            const promoElement = document.createElement('div');
            promoElement.innerHTML = `
                <h4>${promo.ID} - ${promo.Codigo}</h4>
                <p>Desconto: ${promo.Desconto}%</p>
                <p>Produtos: ${promo.Produtos.join(', ')}</p>
                <p>Validade: ${new Date(promo.ValidadeInicio).toLocaleDateString()} - ${new Date(promo.ValidadeFim).toLocaleDateString()}</p>
            `;
            mainContent.appendChild(promoElement);
        });
    }

    function searchPromotions(query) {
        fetchPromotions(); 
    }

    function applyPromotion(codigo, produtos, totalCompra) {
        fetch('http://localhost:8081/promotions/apply', {
            method: 'POST', 
            headers: {
                'Content-Type': 'application/json' 
            },
            body: JSON.stringify({ codigo, produtos, totalCompra }) 
        })
            .then(response => response.json()) 
            .then(data => {
                console.log('Promotion applied:', data); 
                alert(`Desconto: R$${data.desconto}, Total: R$${data.total}`); 
            })
            .catch(error => console.error('Error applying promotion:', error)); 
    }

    applyPromotion('PROMO10', ['produto1', 'produto2'], 100.00);
});
