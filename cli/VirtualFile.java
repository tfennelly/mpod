
public abstract class VirtualFile<T> {

    private Array<String> lines;

    public VirtualFile(String path) {
        // Open and read the file into an array ...
    }

    protected List<String> getLines() {
        return this.lines;
    }
    
    public abstract void printToStdout();
    
    public abstract void sendToOpenTelemetry();
}

public class StatFile extends VirtualFile<List<int>> {
    private Map<String, List<int>> getKVPs() {
        Map<String, List<int>> kvps = new HashMap<String, List<int>>()
        List<String> lines = this.getLines();
        
        for (i = 0; i < lines.size(); i++) {
            String line = lines.get(i);
            StringTokenizer tokenizer = new StringTokenizer(input, " ");

            if (tokenizer.hasMoreTokens()) {
                List<int> values = new Array<int>();
                String lineKey = tokenizer.nextToken();
                
                while (tokenizer.hasMoreTokens()) {
                    String value = tokenizer.nextToken();
                    if (!value.trim().equals(""))) {
                        values.add(Integer.parse(value.trim()));
                    }
                }

                kvps.put(lineKey, values);
            }
        }
        return kvps;
    }

    public void printToStdout() {
        Map<String, List<int>> cpuValues = this.getKVPs();

        for (Map.Entry<String, List<int>> cpuValue : cpuValues.entrySet()) {
            System.out.println("Key: " + cpuValue.getKey() + ", Value: " + cpuValue.getValue());
        }
    }

    public void toOpenTelemetry() {
        Map<String, List<int>> cpuValues = this.getKVPs();

        for (Map.Entry<String, List<int>> cpuValue : cpuValues.entrySet()) {
            // Send to some OTEL api ...
        }
    }
}

public class App {

    public void main() {
        StatFile cpuData = new StatFile("/proc/stat");
        cpuData.printToStdout():
    }
}