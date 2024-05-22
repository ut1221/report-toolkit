package {{ .PackageName }};

import java.util.Date;
import java.math.BigDecimal;

public class {{ .ClassName }} {
{{ range .Fields }}
    private {{ .Type }} {{ .Name }};
{{ end }}
}
