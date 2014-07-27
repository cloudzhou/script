public class Main {

    public static void main(String[] args) {
        Integer[] numarr = new Integer[] { 7, 4, 2, 9 };
        Integer sum = 0;
        for (Integer number : numarr) {
            sum += number;
        }
        if (sum % 2 == 1) {
            System.out.println("NO");
            return;
        }
        Integer halfSum = sum / 2;
        System.out.println(numarrEquals(numarr, 0, halfSum));
    }

    private static Boolean numarrEquals(Integer[] numarr, Integer index, Integer value) {
        if (numarr[index] == value) {
            return true;
        }
        if (index == numarr.length - 1) {
            return false;
        }
        Integer valueIndex = numarr[index];
        index++;
        if (numarrEquals(numarr, index, value - valueIndex)) {
            return true;
        }
        if (numarrEquals(numarr, index, value)) {
            return true;
        }
        return false;
    }
}

