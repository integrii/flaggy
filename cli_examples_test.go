package flaggy_test

import (
	"testing"
	"time"

	"github.com/integrii/flaggy"
)

func TestFlaggyCLIExamples(t *testing.T) {
	tests := []struct {
		name string
		exec func(t *testing.T)
	}{
		{
			name: "AstroScanRoot",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: astro-scan --target "NGC 1300" --exposure 1200 --filters R,G,B
				// Expect: target="NGC 1300", exposure=1200, filters="R,G,B"
				parser := flaggy.NewParser("astro-scan")
				var target string
				var exposure int
				var filters string
				parser.String(&target, "", "target", "")
				parser.Int(&exposure, "", "exposure", "")
				parser.String(&filters, "", "filters", "")
				args := []string{"--target", "NGC 1300", "--exposure", "1200", "--filters", "R,G,B"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if target != "NGC 1300" || exposure != 1200 || filters != "R,G,B" {
					t.Fatalf("unexpected values: target=%q exposure=%d filters=%q", target, exposure, filters)
				}
			},
		},
		{
			name: "AstroScanCalibrate",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: astro-scan calibrate --dark-frames 25 --bias-frames 10 --overwrite
				// Expect: darkFrames=25, biasFrames=10, overwrite=true
				parser := flaggy.NewParser("astro-scan")
				calibrate := flaggy.NewSubcommand("calibrate")
				parser.AttachSubcommand(calibrate, 1)
				var darkFrames, biasFrames int
				var overwrite bool
				calibrate.Int(&darkFrames, "", "dark-frames", "")
				calibrate.Int(&biasFrames, "", "bias-frames", "")
				calibrate.Bool(&overwrite, "", "overwrite", "")
				args := []string{"calibrate", "--dark-frames", "25", "--bias-frames", "10", "--overwrite"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if !calibrate.Used || darkFrames != 25 || biasFrames != 10 || !overwrite {
					t.Fatalf("unexpected values: used=%v dark=%d bias=%d overwrite=%v", calibrate.Used, darkFrames, biasFrames, overwrite)
				}
			},
		},
		{
			name: "AstroScanAnalyzeSpectrum",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: astro-scan analyze spectrum --line-id H-alpha --threshold 0.73 --plot
				// Expect: lineID="H-alpha", threshold=0.73, plot=true
				parser := flaggy.NewParser("astro-scan")
				analyze := flaggy.NewSubcommand("analyze")
				spectrum := flaggy.NewSubcommand("spectrum")
				parser.AttachSubcommand(analyze, 1)
				analyze.AttachSubcommand(spectrum, 1)
				var lineID string
				var threshold float64
				var plot bool
				spectrum.String(&lineID, "", "line-id", "")
				spectrum.Float64(&threshold, "", "threshold", "")
				spectrum.Bool(&plot, "", "plot", "")
				args := []string{"analyze", "spectrum", "--line-id", "H-alpha", "--threshold", "0.73", "--plot"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if !analyze.Used || !spectrum.Used || lineID != "H-alpha" || threshold != 0.73 || !plot {
					t.Fatalf("unexpected values: analyzeUsed=%v spectrumUsed=%v line=%q threshold=%f plot=%v", analyze.Used, spectrum.Used, lineID, threshold, plot)
				}
			},
		},
		{
			name: "DatasmithImport",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: datasmith import --source s3://datasets/raw --schema user_logins --batch-size 5000
				// Expect: source="s3://datasets/raw", schema="user_logins", batchSize=5000
				parser := flaggy.NewParser("datasmith")
				importCmd := flaggy.NewSubcommand("import")
				parser.AttachSubcommand(importCmd, 1)
				var source, schema string
				var batch int
				importCmd.String(&source, "", "source", "")
				importCmd.String(&schema, "", "schema", "")
				importCmd.Int(&batch, "", "batch-size", "")
				args := []string{"import", "--source", "s3://datasets/raw", "--schema", "user_logins", "--batch-size", "5000"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if !importCmd.Used || source != "s3://datasets/raw" || schema != "user_logins" || batch != 5000 {
					t.Fatalf("unexpected values: used=%v source=%q schema=%q batch=%d", importCmd.Used, source, schema, batch)
				}
			},
		},
		{
			name: "DatasmithExportParquet",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: datasmith export parquet --destination /archives/2025 --partition-date 2025-09-19
				// Expect: destination="/archives/2025", partitionDate="2025-09-19"
				parser := flaggy.NewParser("datasmith")
				export := flaggy.NewSubcommand("export")
				parquet := flaggy.NewSubcommand("parquet")
				parser.AttachSubcommand(export, 1)
				export.AttachSubcommand(parquet, 1)
				var destination, partition string
				parquet.String(&destination, "", "destination", "")
				parquet.String(&partition, "", "partition-date", "")
				args := []string{"export", "parquet", "--destination", "/archives/2025", "--partition-date", "2025-09-19"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if !export.Used || !parquet.Used || destination != "/archives/2025" || partition != "2025-09-19" {
					t.Fatalf("unexpected values: exportUsed=%v parquetUsed=%v dest=%q partition=%q", export.Used, parquet.Used, destination, partition)
				}
			},
		},
		{
			name: "DatasmithValidate",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: datasmith validate --rule-set security --rule-set retention --fail-fast
				// Expect: ruleSets=[security retention], failFast=true
				parser := flaggy.NewParser("datasmith")
				validate := flaggy.NewSubcommand("validate")
				parser.AttachSubcommand(validate, 1)
				var ruleSets []string
				var failFast bool
				validate.StringSlice(&ruleSets, "", "rule-set", "")
				validate.Bool(&failFast, "", "fail-fast", "")
				args := []string{"validate", "--rule-set", "security", "--rule-set", "retention", "--fail-fast"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if len(ruleSets) != 2 || ruleSets[0] != "security" || ruleSets[1] != "retention" || !failFast {
					t.Fatalf("unexpected values: rules=%v failFast=%v", ruleSets, failFast)
				}
			},
		},
		{
			name: "FarmFieldStatus",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: farmctl field status --field-id 12 --verbose
				// Expect: fieldID=12, verbose=true
				parser := flaggy.NewParser("farmctl")
				field := flaggy.NewSubcommand("field")
				status := flaggy.NewSubcommand("status")
				parser.AttachSubcommand(field, 1)
				field.AttachSubcommand(status, 1)
				var fieldID int
				var verbose bool
				status.Int(&fieldID, "", "field-id", "")
				status.Bool(&verbose, "", "verbose", "")
				args := []string{"field", "status", "--field-id", "12", "--verbose"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if !field.Used || !status.Used || fieldID != 12 || !verbose {
					t.Fatalf("unexpected values: fieldUsed=%v statusUsed=%v id=%d verbose=%v", field.Used, status.Used, fieldID, verbose)
				}
			},
		},
		{
			name: "FarmIrrigationSchedule",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: farmctl irrigation schedule --field-id 12 --start 05:00 --duration 30m --days mon,wed,fri
				// Expect: fieldID=12, start="05:00", duration=30m, days="mon,wed,fri"
				parser := flaggy.NewParser("farmctl")
				irrigation := flaggy.NewSubcommand("irrigation")
				schedule := flaggy.NewSubcommand("schedule")
				parser.AttachSubcommand(irrigation, 1)
				irrigation.AttachSubcommand(schedule, 1)
				var fieldID int
				var start string
				var duration time.Duration
				var days string
				schedule.Int(&fieldID, "", "field-id", "")
				schedule.String(&start, "", "start", "")
				schedule.Duration(&duration, "", "duration", "")
				schedule.String(&days, "", "days", "")
				args := []string{"irrigation", "schedule", "--field-id", "12", "--start", "05:00", "--duration", "30m", "--days", "mon,wed,fri"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				expectDur, _ := time.ParseDuration("30m")
				if fieldID != 12 || start != "05:00" || duration != expectDur || days != "mon,wed,fri" {
					t.Fatalf("unexpected values: id=%d start=%q duration=%v days=%q", fieldID, start, duration, days)
				}
			},
		},
		{
			name: "FarmDroneLaunchScout",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: farmctl drone launch scout --altitude 120 --speed 15 --waypoints north,center,south
				// Expect: altitude=120, speed=15, waypoints="north,center,south"
				parser := flaggy.NewParser("farmctl")
				drone := flaggy.NewSubcommand("drone")
				launch := flaggy.NewSubcommand("launch")
				scout := flaggy.NewSubcommand("scout")
				parser.AttachSubcommand(drone, 1)
				drone.AttachSubcommand(launch, 1)
				launch.AttachSubcommand(scout, 1)
				var altitude, speed int
				var waypoints string
				scout.Int(&altitude, "", "altitude", "")
				scout.Int(&speed, "", "speed", "")
				scout.String(&waypoints, "", "waypoints", "")
				args := []string{"drone", "launch", "scout", "--altitude", "120", "--speed", "15", "--waypoints", "north,center,south"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if altitude != 120 || speed != 15 || waypoints != "north,center,south" {
					t.Fatalf("unexpected values: altitude=%d speed=%d waypoints=%q", altitude, speed, waypoints)
				}
			},
		},
		{
			name: "GalaxyDeployApi",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: galaxyctl deploy api --env staging -r v2.1.4 --rollback
				// Expect: env="staging", release="v2.1.4", rollback=true
				parser := flaggy.NewParser("galaxyctl")
				deploy := flaggy.NewSubcommand("deploy")
				api := flaggy.NewSubcommand("api")
				parser.AttachSubcommand(deploy, 1)
				deploy.AttachSubcommand(api, 1)
				var env, release string
				var rollback bool
				api.String(&env, "", "env", "")
				api.String(&release, "r", "release", "")
				api.Bool(&rollback, "", "rollback", "")
				args := []string{"deploy", "api", "--env", "staging", "-r", "v2.1.4", "--rollback"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if env != "staging" || release != "v2.1.4" || !rollback {
					t.Fatalf("unexpected values: env=%q release=%q rollback=%v", env, release, rollback)
				}
			},
		},
		{
			name: "GalaxyDeployWorker",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: galaxyctl deploy worker --env prod --config configs/worker.yaml --wait
				// Expect: env="prod", config="configs/worker.yaml", wait=true
				parser := flaggy.NewParser("galaxyctl")
				deploy := flaggy.NewSubcommand("deploy")
				worker := flaggy.NewSubcommand("worker")
				parser.AttachSubcommand(deploy, 1)
				deploy.AttachSubcommand(worker, 1)
				var env, config string
				var wait bool
				worker.String(&env, "", "env", "")
				worker.String(&config, "", "config", "")
				worker.Bool(&wait, "", "wait", "")
				args := []string{"deploy", "worker", "--env", "prod", "--config", "configs/worker.yaml", "--wait"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if env != "prod" || config != "configs/worker.yaml" || !wait {
					t.Fatalf("unexpected values: env=%q config=%q wait=%v", env, config, wait)
				}
			},
		},
		{
			name: "GalaxyLogsWorker",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: galaxyctl logs worker --env prod --since 4h --tail 200 --color
				// Expect: env="prod", since=4h, tail=200, color=true
				parser := flaggy.NewParser("galaxyctl")
				logs := flaggy.NewSubcommand("logs")
				worker := flaggy.NewSubcommand("worker")
				parser.AttachSubcommand(logs, 1)
				logs.AttachSubcommand(worker, 1)
				var env string
				var since time.Duration
				var tail int
				var color bool
				worker.String(&env, "", "env", "")
				worker.Duration(&since, "", "since", "")
				worker.Int(&tail, "", "tail", "")
				worker.Bool(&color, "", "color", "")
				args := []string{"logs", "worker", "--env", "prod", "--since", "4h", "--tail", "200", "--color"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				expect, _ := time.ParseDuration("4h")
				if env != "prod" || since != expect || tail != 200 || !color {
					t.Fatalf("unexpected values: env=%q since=%v tail=%d color=%v", env, since, tail, color)
				}
			},
		},
		{
			name: "HoloeditRenderStoryboard",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: holoedit render storyboard --resolution 4k --fps 60 --output ./renders/scene01
				// Expect: resolution="4k", fps=60, output="./renders/scene01"
				parser := flaggy.NewParser("holoedit")
				render := flaggy.NewSubcommand("render")
				storyboard := flaggy.NewSubcommand("storyboard")
				parser.AttachSubcommand(render, 1)
				render.AttachSubcommand(storyboard, 1)
				var resolution string
				var fps int
				var output string
				storyboard.String(&resolution, "", "resolution", "")
				storyboard.Int(&fps, "", "fps", "")
				storyboard.String(&output, "", "output", "")
				args := []string{"render", "storyboard", "--resolution", "4k", "--fps", "60", "--output", "./renders/scene01"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if resolution != "4k" || fps != 60 || output != "./renders/scene01" {
					t.Fatalf("unexpected values: resolution=%q fps=%d output=%q", resolution, fps, output)
				}
			},
		},
		{
			name: "HoloeditPublish",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: holoedit publish --channel beta --notes "Testing volumetric transitions" --tag sfx
				// Expect: channel="beta", notes="Testing volumetric transitions", tag="sfx"
				parser := flaggy.NewParser("holoedit")
				publish := flaggy.NewSubcommand("publish")
				parser.AttachSubcommand(publish, 1)
				var channel, notes, tag string
				publish.String(&channel, "", "channel", "")
				publish.String(&notes, "", "notes", "")
				publish.String(&tag, "", "tag", "")
				args := []string{"publish", "--channel", "beta", "--notes", "Testing volumetric transitions", "--tag", "sfx"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if !publish.Used || channel != "beta" || notes != "Testing volumetric transitions" || tag != "sfx" {
					t.Fatalf("unexpected values: used=%v channel=%q notes=%q tag=%q", publish.Used, channel, notes, tag)
				}
			},
		},
		{
			name: "HoloeditAssetsSync",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: holoedit assets sync --source ./assets --dest s3://holoedit-assets --include textures --exclude temp
				// Expect: source="./assets", dest="s3://holoedit-assets", include="textures", exclude="temp"
				parser := flaggy.NewParser("holoedit")
				assets := flaggy.NewSubcommand("assets")
				sync := flaggy.NewSubcommand("sync")
				parser.AttachSubcommand(assets, 1)
				assets.AttachSubcommand(sync, 1)
				var source, dest, include, exclude string
				sync.String(&source, "", "source", "")
				sync.String(&dest, "", "dest", "")
				sync.String(&include, "", "include", "")
				sync.String(&exclude, "", "exclude", "")
				args := []string{"assets", "sync", "--source", "./assets", "--dest", "s3://holoedit-assets", "--include", "textures", "--exclude", "temp"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if source != "./assets" || dest != "s3://holoedit-assets" || include != "textures" || exclude != "temp" {
					t.Fatalf("unexpected values: source=%q dest=%q include=%q exclude=%q", source, dest, include, exclude)
				}
			},
		},
		{
			name: "NanodevBuildSensor",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: nanodev build sensor --board nanoX -O2 --define DEBUG=0 --flash
				// Expect: board="nanoX", optLevel2=true, define="DEBUG=0", flash=true
				parser := flaggy.NewParser("nanodev")
				build := flaggy.NewSubcommand("build")
				sensor := flaggy.NewSubcommand("sensor")
				parser.AttachSubcommand(build, 1)
				build.AttachSubcommand(sensor, 1)
				var board, define string
				var optLevel2, flash bool
				sensor.String(&board, "", "board", "")
				sensor.Bool(&optLevel2, "O2", "optimize-2", "")
				sensor.String(&define, "", "define", "")
				sensor.Bool(&flash, "", "flash", "")
				args := []string{"build", "sensor", "--board", "nanoX", "-O2", "--define", "DEBUG=0", "--flash"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if board != "nanoX" || !optLevel2 || define != "DEBUG=0" || !flash {
					t.Fatalf("unexpected values: board=%q opt=%v define=%q flash=%v", board, optLevel2, define, flash)
				}
			},
		},
		{
			name: "NanodevTestSensorRegression",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: nanodev test sensor regression --pattern smoke --repeat 5 --seed 42
				// Expect: pattern="smoke", repeat=5, seed=42
				parser := flaggy.NewParser("nanodev")
				testCmd := flaggy.NewSubcommand("test")
				sensor := flaggy.NewSubcommand("sensor")
				regression := flaggy.NewSubcommand("regression")
				parser.AttachSubcommand(testCmd, 1)
				testCmd.AttachSubcommand(sensor, 1)
				sensor.AttachSubcommand(regression, 1)
				var pattern string
				var repeat, seed int
				regression.String(&pattern, "", "pattern", "")
				regression.Int(&repeat, "", "repeat", "")
				regression.Int(&seed, "", "seed", "")
				args := []string{"test", "sensor", "regression", "--pattern", "smoke", "--repeat", "5", "--seed", "42"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if pattern != "smoke" || repeat != 5 || seed != 42 {
					t.Fatalf("unexpected values: pattern=%q repeat=%d seed=%d", pattern, repeat, seed)
				}
			},
		},
		{
			name: "NanodevMonitor",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: nanodev monitor --port /dev/ttyUSB0 --baud 115200 --raw
				// Expect: port="/dev/ttyUSB0", baud=115200, raw=true
				parser := flaggy.NewParser("nanodev")
				monitor := flaggy.NewSubcommand("monitor")
				parser.AttachSubcommand(monitor, 1)
				var port string
				var baud int
				var raw bool
				monitor.String(&port, "", "port", "")
				monitor.Int(&baud, "", "baud", "")
				monitor.Bool(&raw, "", "raw", "")
				args := []string{"monitor", "--port", "/dev/ttyUSB0", "--baud", "115200", "--raw"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if !monitor.Used || port != "/dev/ttyUSB0" || baud != 115200 || !raw {
					t.Fatalf("unexpected values: used=%v port=%q baud=%d raw=%v", monitor.Used, port, baud, raw)
				}
			},
		},
		{
			name: "QuantifySimulatePortfolioRebalance",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: quantify simulate portfolio rebalance --strategy momentum --samples 10000 --threads 8
				// Expect: strategy="momentum", samples=10000, threads=8
				parser := flaggy.NewParser("quantify")
				simulate := flaggy.NewSubcommand("simulate")
				portfolio := flaggy.NewSubcommand("portfolio")
				rebalance := flaggy.NewSubcommand("rebalance")
				parser.AttachSubcommand(simulate, 1)
				simulate.AttachSubcommand(portfolio, 1)
				portfolio.AttachSubcommand(rebalance, 1)
				var strategy string
				var samples, threads int
				rebalance.String(&strategy, "", "strategy", "")
				rebalance.Int(&samples, "", "samples", "")
				rebalance.Int(&threads, "", "threads", "")
				args := []string{"simulate", "portfolio", "rebalance", "--strategy", "momentum", "--samples", "10000", "--threads", "8"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if strategy != "momentum" || samples != 10000 || threads != 8 {
					t.Fatalf("unexpected values: strategy=%q samples=%d threads=%d", strategy, samples, threads)
				}
			},
		},
		{
			name: "QuantifyBacktestEquities",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: quantify backtest equities --from 2015-01-01 --to 2024-12-31 --fee-model tiered --report summary
				// Expect: from="2015-01-01", to="2024-12-31", feeModel="tiered", report="summary"
				parser := flaggy.NewParser("quantify")
				backtest := flaggy.NewSubcommand("backtest")
				equities := flaggy.NewSubcommand("equities")
				parser.AttachSubcommand(backtest, 1)
				backtest.AttachSubcommand(equities, 1)
				var from, to, feeModel, report string
				equities.String(&from, "", "from", "")
				equities.String(&to, "", "to", "")
				equities.String(&feeModel, "", "fee-model", "")
				equities.String(&report, "", "report", "")
				args := []string{"backtest", "equities", "--from", "2015-01-01", "--to", "2024-12-31", "--fee-model", "tiered", "--report", "summary"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if from != "2015-01-01" || to != "2024-12-31" || feeModel != "tiered" || report != "summary" {
					t.Fatalf("unexpected values: from=%q to=%q feeModel=%q report=%q", from, to, feeModel, report)
				}
			},
		},
		{
			name: "QuantifyRiskStress",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: quantify risk stress --scenario "usd+200bps" --scenario "oil-30%" --parallel
				// Expect: scenarios=["usd+200bps", "oil-30%"], parallel=true
				parser := flaggy.NewParser("quantify")
				risk := flaggy.NewSubcommand("risk")
				stress := flaggy.NewSubcommand("stress")
				parser.AttachSubcommand(risk, 1)
				risk.AttachSubcommand(stress, 1)
				var scenarios []string
				var parallel bool
				stress.StringSlice(&scenarios, "", "scenario", "")
				stress.Bool(&parallel, "", "parallel", "")
				args := []string{"risk", "stress", "--scenario", "usd+200bps", "--scenario", "oil-30%", "--parallel"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if len(scenarios) != 2 || scenarios[0] != "usd+200bps" || scenarios[1] != "oil-30%" || !parallel {
					t.Fatalf("unexpected values: scenarios=%v parallel=%v", scenarios, parallel)
				}
			},
		},
		{
			name: "TerraforgePlanCityBuild",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: terraforge plan city build --map newhaven --population 850000 --zoning mixed
				// Expect: map="newhaven", population=850000, zoning="mixed"
				parser := flaggy.NewParser("terraforge")
				plan := flaggy.NewSubcommand("plan")
				city := flaggy.NewSubcommand("city")
				build := flaggy.NewSubcommand("build")
				parser.AttachSubcommand(plan, 1)
				plan.AttachSubcommand(city, 1)
				city.AttachSubcommand(build, 1)
				var mapName, zoning string
				var population int
				build.String(&mapName, "", "map", "")
				build.Int(&population, "", "population", "")
				build.String(&zoning, "", "zoning", "")
				args := []string{"plan", "city", "build", "--map", "newhaven", "--population", "850000", "--zoning", "mixed"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if mapName != "newhaven" || population != 850000 || zoning != "mixed" {
					t.Fatalf("unexpected values: map=%q population=%d zoning=%q", mapName, population, zoning)
				}
			},
		},
		{
			name: "TerraforgeApplyCityBuild",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: terraforge apply city build --dry-run --state ./statefiles/newhaven.tfstate
				// Expect: dryRun=true, state="./statefiles/newhaven.tfstate"
				parser := flaggy.NewParser("terraforge")
				apply := flaggy.NewSubcommand("apply")
				city := flaggy.NewSubcommand("city")
				build := flaggy.NewSubcommand("build")
				parser.AttachSubcommand(apply, 1)
				apply.AttachSubcommand(city, 1)
				city.AttachSubcommand(build, 1)
				var dryRun bool
				var state string
				build.Bool(&dryRun, "", "dry-run", "")
				build.String(&state, "", "state", "")
				args := []string{"apply", "city", "build", "--dry-run", "--state", "./statefiles/newhaven.tfstate"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if !dryRun || state != "./statefiles/newhaven.tfstate" {
					t.Fatalf("unexpected values: dryRun=%v state=%q", dryRun, state)
				}
			},
		},
		{
			name: "TerraforgeInspectNetwork",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: terraforge inspect network --layer transit --format json --pretty --limit 50
				// Expect: layer="transit", format="json", pretty=true, limit=50
				parser := flaggy.NewParser("terraforge")
				inspect := flaggy.NewSubcommand("inspect")
				network := flaggy.NewSubcommand("network")
				parser.AttachSubcommand(inspect, 1)
				inspect.AttachSubcommand(network, 1)
				var layer, format string
				var pretty bool
				var limit int
				network.String(&layer, "", "layer", "")
				network.String(&format, "", "format", "")
				network.Bool(&pretty, "", "pretty", "")
				network.Int(&limit, "", "limit", "")
				args := []string{"inspect", "network", "--layer", "transit", "--format", "json", "--pretty", "--limit", "50"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if layer != "transit" || format != "json" || !pretty || limit != 50 {
					t.Fatalf("unexpected values: layer=%q format=%q pretty=%v limit=%d", layer, format, pretty, limit)
				}
			},
		},
		{
			name: "TrailmixIngestTelemetry",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: trailmix ingest telemetry --input ./logs/*.json --profile production -- --follow --color
				// Expect: input="./logs/*.json", profile="production", trailing=[--follow --color]
				parser := flaggy.NewParser("trailmix")
				ingest := flaggy.NewSubcommand("ingest")
				telemetry := flaggy.NewSubcommand("telemetry")
				parser.AttachSubcommand(ingest, 1)
				ingest.AttachSubcommand(telemetry, 1)
				var input, profile string
				telemetry.String(&input, "", "input", "")
				telemetry.String(&profile, "", "profile", "")
				args := []string{"ingest", "telemetry", "--input", "./logs/*.json", "--profile", "production", "--", "--follow", "--color"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				if input != "./logs/*.json" || profile != "production" {
					t.Fatalf("unexpected values: input=%q profile=%q", input, profile)
				}
				if len(parser.TrailingArguments) != 2 || parser.TrailingArguments[0] != "--follow" || parser.TrailingArguments[1] != "--color" {
					t.Fatalf("unexpected trailing: %v", parser.TrailingArguments)
				}
			},
		},
		{
			name: "TrailmixQueryMetricsLatency",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: trailmix query metrics latency --window 15m --aggregate p99 --group-by endpoint
				// Expect: window=15m, aggregate="p99", groupBy="endpoint"
				parser := flaggy.NewParser("trailmix")
				query := flaggy.NewSubcommand("query")
				metrics := flaggy.NewSubcommand("metrics")
				latency := flaggy.NewSubcommand("latency")
				parser.AttachSubcommand(query, 1)
				query.AttachSubcommand(metrics, 1)
				metrics.AttachSubcommand(latency, 1)
				var window time.Duration
				var aggregate, groupBy string
				latency.Duration(&window, "", "window", "")
				latency.String(&aggregate, "", "aggregate", "")
				latency.String(&groupBy, "", "group-by", "")
				args := []string{"query", "metrics", "latency", "--window", "15m", "--aggregate", "p99", "--group-by", "endpoint"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				expect, _ := time.ParseDuration("15m")
				if window != expect || aggregate != "p99" || groupBy != "endpoint" {
					t.Fatalf("unexpected values: window=%v aggregate=%q groupBy=%q", window, aggregate, groupBy)
				}
			},
		},
		{
			name: "TrailmixAlertCreate",
			exec: func(t *testing.T) {
				t.Parallel()
				// Command: trailmix alert create --name "CPU Spike" --threshold 85 --duration 10m --notify pagerduty
				// Expect: name="CPU Spike", threshold=85, duration=10m, notify="pagerduty"
				parser := flaggy.NewParser("trailmix")
				alert := flaggy.NewSubcommand("alert")
				create := flaggy.NewSubcommand("create")
				parser.AttachSubcommand(alert, 1)
				alert.AttachSubcommand(create, 1)
				var name, notify string
				var threshold int
				var duration time.Duration
				create.String(&name, "", "name", "")
				create.Int(&threshold, "", "threshold", "")
				create.Duration(&duration, "", "duration", "")
				create.String(&notify, "", "notify", "")
				args := []string{"alert", "create", "--name", "CPU Spike", "--threshold", "85", "--duration", "10m", "--notify", "pagerduty"}
				if err := parser.ParseArgs(args); err != nil {
					t.Fatalf("parse error: %v", err)
				}
				expect, _ := time.ParseDuration("10m")
				if name != "CPU Spike" || threshold != 85 || duration != expect || notify != "pagerduty" {
					t.Fatalf("unexpected values: name=%q threshold=%d duration=%v notify=%q", name, threshold, duration, notify)
				}
			},
		},
	}

	for _, tc := range tests {
		c := tc
		t.Run(c.name, c.exec)
	}
}
