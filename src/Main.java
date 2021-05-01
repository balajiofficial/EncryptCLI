import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) throws IOException {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter 1 for encryption or 2 for decryption : ");
        String num = scanner.nextLine();
        String key = generateRandomKey();
        if (num.equals("1")) {
            System.out.print("Enter the name of the file to be encrypted : ");
            String file = scanner.nextLine();
            FileReader reader = null;
            try {
                reader = new FileReader(file);
            } catch (FileNotFoundException e) {
                System.out.println("Please enter the name of a valid file.");
                System.out.println("Encryption Aborted.");
                System.exit(-1);
            }
            int data = reader.read();
            StringBuilder str = new StringBuilder();
            while (data != -1) {
                str.append(data).append(" ");
                data = reader.read();
            }
            String newFile = "EncryptCLI." + file.substring(file.indexOf('.') + 1);
            reader.close();
            FileWriter writer = new FileWriter(newFile);
            writer.write(str.toString());
            writer.close();
            System.out.println("Encrypted at " + newFile);
        } else if (num.equals("2")) {
            System.out.print("Enter the name of the file to be decrypted : ");
            String file = scanner.nextLine();
            FileReader reader = null;
            try {
                reader = new FileReader(file);
            } catch (FileNotFoundException e) {
                System.out.println("Please enter the name of a valid file.");
                System.out.println("Decryption Aborted.");
                System.exit(-1);
            }
            int temp = reader.read();
            char data = (char)temp;
            StringBuilder str = new StringBuilder();
            StringBuilder te = new StringBuilder();
            while(temp != -1) {
                if (data == ' ') {
                    str.append((char) Integer.parseInt(te.toString()));
                    te = new StringBuilder();
                } else {
                    te.append(data);
                }
                temp = reader.read();
                data = (char)temp;
            }
            str.append(te);
            String newFile = "DecryptCLI." + file.substring(file.indexOf('.') + 1);
            FileWriter writer = new FileWriter(newFile);
            writer.write(str.toString());
            writer.close();
            reader.close();
            System.out.println("Decrypted at " + newFile);
        }
    }

    public static String generateRandomKey() {
        StringBuilder key = new StringBuilder();
        for (int i = 0; i < 15; ++i) {
            key.append(Math.round(Math.random() * 93 + 33));
        }
        return key.toString();
    }
}
