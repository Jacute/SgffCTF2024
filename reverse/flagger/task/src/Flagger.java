import java.util.Scanner;

public class Flagger {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Input flag: ");
        String flag = scanner.nextLine();
        scanner.close();

        if (flag.length() != 32) {
            System.out.println("Incorrect flag length");
            System.exit(1);
        } else if (!flag.endsWith("p}")) {
            System.out.println("Incorrect flag 1");
            System.exit(1);
        } else if (!flag.startsWith("SgffCTF{")) {
            System.out.println("Incorrect flag 2");
            System.exit(1);
        } else if (!flag.substring(27, 30).equals("u_u")) {
            System.out.println("Incorrect flag 3");
            System.exit(1);
        } else if (!flag.substring(12, 17).equals("6_55r")) {
            System.out.println("Incorrect flag 4");
            System.exit(1);
        } else if (!flag.substring(17, 23).equals("12gf11")) {
            System.out.println("Incorrect flag 5");
            System.exit(1);
        } else if (!flag.substring(23, 27).equals("5_pp")) {
            System.out.println("Incorrect flag 6");
            System.exit(1);
        } else if (!flag.substring(8, 12).equals("pvpv")) {
            System.out.println("Incorrect flag 7");
            System.exit(1);
        }
        System.out.println("Yeah! Your flag is " + flag);
    }
}