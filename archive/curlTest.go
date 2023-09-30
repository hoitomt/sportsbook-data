package archive

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

var cookiePath string

func init() {
	cookiePath = os.Getenv("COOKIE_PATH")
	if len(cookiePath) == 0 {
		exePath, err := os.Executable()
		if err != nil {
			log.Fatal().Msgf("Error retrieving the currently running directory %s", err)
		}
		exeDir := filepath.Dir(exePath)
		log.Info().Msgf("Running in: %s", exeDir)

		cookieDir := filepath.Join(exeDir, "cookies")
		err = os.MkdirAll(cookieDir, os.ModePerm)
		if err != nil {
			log.Fatal().Msgf("Error creating the cookies directory: %s", err)
		}
		cookiePath = filepath.Join(cookieDir, "sb_cookies.txt")
	}
	if err := os.Remove(cookiePath); err != nil {
		log.Fatal().Msgf("Error creating the cookiefile %s", err)
	}
}

func main() {
	loginPostUrl := "https://www.sportsbook.ag/cca/customerauthn/pl/login"
	username := os.Getenv("SPORTSBOOK_USERNAME")
	password := os.Getenv("SPORTSBOOK_PASSWORD")
	blackbox := os.Getenv("SPORTSBOOK_BLACKBOX")

	loginCmdAttr := ""
	// loginCmdAttr := fmt.Sprintf("--cookie %s", cookiePath)
	// loginCmdAttr += fmt.Sprintf(" --cookie-jar %s", cookiePath)
	// loginCmdAttr += fmt.Sprintf("-d username=%s", username)
	// loginCmdAttr += fmt.Sprintf(" -d password=%s", password)
	// loginCmdAttr += fmt.Sprintf(" -d service=%s", "https://www.sportsbook.ag/livesports")
	// loginCmdAttr += fmt.Sprintf(" -d login_fail=%s", "https://www.sportsbook.ag/login?service=livesports")
	// loginCmdAttr += " -d sp_casinoid=7000"
	// loginCmdAttr += " -d sp_siteID=7000"
	// loginCmdAttr += " -d lp_casinoid=7000"
	// loginCmdAttr += fmt.Sprintf(" -d blackbox=%s", blackbox)
	// loginCmdAttr += fmt.Sprintf(" %s", loginPostUrl)

	// secChUa := "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Google Chrome\";v=\"116\""

	log.Info().Msgf("Login cmd: curl %s", loginCmdAttr)

	curlCmd := exec.Command("curl",
		fmt.Sprintf("-b %s", cookiePath),
		fmt.Sprintf("-c %s", cookiePath),
		"-L",
		"-s",
		fmt.Sprintf("-d username=%s", username),
		fmt.Sprintf("-d password=%s", password),
		fmt.Sprintf("-d service=%s", "https://www.sportsbook.ag/livesports"),
		fmt.Sprintf("-d login_fail=%s", "https://www.sportsbook.ag/login?service=livesports"),
		fmt.Sprintf(" -d blackbox=%s", blackbox),
		"-d sp_casinoid=7000",
		"-d sp_siteID=7000",
		"-d lp_casinoid=7000",
		loginPostUrl,
	)

	log.Info().Msgf("%#v", curlCmd.Args)

	output, err := curlCmd.CombinedOutput()
	if err != nil {
		log.Fatal().Msgf("Error running the curl command: %s", err)
	}

	log.Info().Msgf("OUTPUT: %s", string(output))
}

// <<-EOF
//   curl --cookie #{cookie_path} --cookie-jar #{cookie_path} -L -s \
// -d "username=#{username}" \
//    -d "password=#{password}" \
//    -d "service=https://www.sportsbook.ag/livesports" \
//    -d "login_fail=https://www.sportsbook.ag/login?service=livesports" \
//    -d "sp_casinoid=7000" \
//    -d "sp_siteID=7000" \
//    -d "lp_casinoid=7000" \
//    -d "blackbox=#{blackbox}" #{Rails.configuration.SB_LOGIN_URL}
//   EOF
