cd /Users/zhouyaoxu/Desktop/参考设计
blastn -query 细蜉科.fasta -db TAIR10 -evalue 1e-6 -outfmt 6 -num_threads 6 -out out_file
echo "Blast done"