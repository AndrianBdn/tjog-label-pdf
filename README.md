# IKEA TJOG Label PDF 

This repository contains the PDF template for IKEA TJOG labels, and the Go code to generate the PDF files.

IKEA TJOG is a paperboard [storage box](https://www.ikea.com/ae/en/p/tjog-storage-box-with-lid-dark-grey-40477665/) / [magazine file](https://www.ikea.com/ae/en/p/tjog-magazine-file-dark-grey-90477658/).

The labels are 51mmx44mm, supposed to be folded in two and inserted in the metallic table holder. The IKEA assembly instructions contains one label that could be cut out. However sometimes you need to replace the labels, and hence this repository. 

You can generate the labels with the go code, or use PDF files [tjog_labels_A4.pdf](tjog_labels_A4.pdf) or [tjog_labels_Letter.pdf](tjog_labels_Letter.pdf) directly. 

You can add inscriptions to the labels by annotating the file in any PDF viewer software (use the top half of each label, above the fold line).

## Printing Note

When printing the resulting files make sure to select 100% scale (fit to page may not work well even when printing to the same page size).

## Usage 

1. Download either [tjog-labels-A4.pdf](tjog-labels-A4.pdf) or [tjog-labels-Letter.pdf](tjog-labels-Letter.pdf)
2. Cut out the labels (black lines indicate cutting lines, gray lines indicate folding)
3. Fold along the gray lines
4. Insert into the metallic label holder

![tjog-label](https://github.com/user-attachments/assets/9f46bec9-4316-4cc9-9ee7-825babf7db6a)

See [IKEA TJOG Assembly Instructions](https://www.ikea.com/ca/en/assembly_instructions/tjog-magazine-file-dark-gray__AA-2201975-1-2.pdf) page 4 for more details.

## Development

To generate the labels yourself, run:

```bash
go run main.go
```

Pre-generated PDF files are already committed to this repository for convenience.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

