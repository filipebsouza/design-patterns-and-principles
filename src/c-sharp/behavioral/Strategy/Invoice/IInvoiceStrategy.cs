namespace Strategy.Invoice
{
    public interface IInvoiceStrategy
    {
        void Generate(Order order);
    }
}
