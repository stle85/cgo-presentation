package profiler

import (
	"fmt"
	"os"
	"os/exec"
)

// convertToSvg converts pprof file to SVG image.
func convertToSvg(pprofPath string) error {
	cmd := exec.Command("go", "tool", "pprof", "--svg", pprofPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error during convert file '%s': %w", pprofPath, err)
	}
	svgPath := pprofPath[0:len(pprofPath)-len("pprof")] + "svg"
	var file *os.File
	file, err = os.OpenFile(svgPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		return fmt.Errorf("cannot open output file '%s': %w", svgPath, err)
	}
	_, err = file.Write(out)
	if err != nil {
		return fmt.Errorf("cannot write to output file '%s': %w", svgPath, err)
	}
	err = file.Close()
	if err != nil {
		return fmt.Errorf("cannot close to output file '%s': %w", svgPath, err)
	}
	return nil
}
