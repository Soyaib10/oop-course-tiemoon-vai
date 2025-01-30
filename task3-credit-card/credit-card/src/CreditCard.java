public class CreditCard {
    private String name;
    private double balance;
    private double dailyWithdrawal;
    private double totalSpent;

    public static final double MAX_CREDIT = 500000;

    public CreditCard(String name) {
        this.name = name;
        this.balance = MAX_CREDIT;
        this.dailyWithdrawal = 0;
        this.totalSpent = 0;
    }

    public void withdraw(double amount) {
        if (amount <= 0 || amount > 20000) {
            throw new RuntimeException("Withdrawal amount must be less than or equal to 20000");
        }

        if (dailyWithdrawal + amount > 100000) {
            throw new RuntimeException("Daily withdrawal limit exceeded");
        }

        if (balance - amount < 0) {
            throw new RuntimeException("Insufficient balance");
        }
        dailyWithdrawal += amount;
        balance -= amount;
    }

    public void payBill(double amount) {
        if (totalSpent + amount > MAX_CREDIT) {
            throw new RuntimeException("Spending limit exceeded");
        }
        totalSpent += amount;
        balance -= amount;
    }

    public void display() {
        System.out.println("Customer name: " + name);
        System.out.println("Available Balance: " + balance);
        System.out.println("Total Spent: " + totalSpent);
        System.out.println("Daily Withdrawal: " + dailyWithdrawal);
    }
}
