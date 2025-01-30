public class BankAccount {
    private String accountNumber;
    private String accountName;
    private double balance;

    public BankAccount(String accountNumber, String accountName, double balance) {
        this.accountNumber = accountNumber;
        this.accountName = accountName;
        this.balance = balance;
    }

    // Deposit method
    public void deposit(double amount) {
        try {
            if (amount <= 0) {
                throw new Exception("Deposit amount must be positive");
            }
            balance += amount;
            System.out.println("Deposit successful");
        } catch (Exception e) {
            System.out.println("Error in deposit: " + e.getMessage());
        }
    }

    // Withdraw method
    public void withdraw(double amount) {
        try {
            if (amount <= 0 || amount > balance) {
                throw new Exception("Invalid withdrawal amount");
            }
            balance -= amount;
            System.out.println("Withdrawal successful");
        } catch (Exception e) {
            System.out.println("Error in withdrawal: " + e.getMessage());
        }
    }

    // Transfer method
    public void transfer(BankAccount receiver, double amount) {
        try {
            if (amount <= 0 || amount > balance) {
                throw new Exception("Invalid transfer amount");
            }
            balance -= amount;
            receiver.balance += amount;
            System.out.println("Transfer successful");
        } catch (Exception e) {
            System.out.println("Error in transfer: " + e.getMessage());
        }
    }

    // Display account info
    public void display() {
        System.out.println("Account Number: " + accountNumber);
        System.out.println("Account Name: " + accountName);
        System.out.println("Balance: " + balance);
    }
}
