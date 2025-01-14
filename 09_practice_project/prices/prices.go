package prices

import (
	"example/practice_project/conversion"
	"example/practice_project/filemanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrice        []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrice {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteJson(job)
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLine()
	if err != nil {
		fmt.Println("Error reading prices.txt:", err)
		return
	}
	fixedPrice, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println("Error converting prices:", err)

		return
	}

	job.InputPrice = fixedPrice

}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,

		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}
