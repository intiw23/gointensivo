package database

import (
	"database/sql"
	"testing"

	"github.com/intiw23/gointensivo/internal/order/entity"
	"github.com/stretchr/testify/suite"

	// link do sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

// executado sempre que é chamado uma suite de testes
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	// abre conexão com banco
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	// cria tabela orders
	// o Go obriga usar todas as variaveis declaradas!
	// Então, esse underline "_" é colocado para evitar o checkup do numeros de regitros inseridos ...para este caso para ignorar.
	_, err = db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	// fecha conexão com banco
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	// nova ordem
	order, err := entity.NewOrder("123", 10.0, 2.0)
	// não pode ter erros
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())

	// injeção da conexão com o banco
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order // criado em branco (orderResult). Depóis vai ser prenchida com o valor que bate no banco de dados(scan)
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
