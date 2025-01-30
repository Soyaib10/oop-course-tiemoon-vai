public class Main {
    public static void main(String[] args) {
        BankAccount soyaib = new BankAccount("123456", "Alice", 5000);
        BankAccount zihad = new BankAccount("789012", "Bob", 3000);

        soyaib.deposit(-100);
        soyaib.withdraw(6000);
        soyaib.transfer(zihad, 1000);

        soyaib.display();
        soyaib.display();
    }
}
