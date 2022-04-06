package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository}
}

func (u *TransactionRepository) Pay(amount int) (int, error) {
	totalPrize, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}

	// not required
	// if totalPrize > amount {
	// 	return 0, fmt.Errorf("insufficient amount")
	// }

	return amount - totalPrize, nil
}
