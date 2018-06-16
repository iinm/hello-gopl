#!/usr/bin/env bash


tmp_dir=$(mktemp -d)
trap "echo -e '\n# cleanup tmp dir'; rm -rfv $tmp_dir" EXIT

if [ ! -d "$tmp_dir" ]; then
  echo "failed to create tmp dir"
  exit 1
fi


input1=$(cat <<EOF
foo
bar
foo
baz
baz
EOF
)

input2=$(cat <<EOF
foo
bar
bar
baz
EOF
)

# setup test data
echo "$input1" > $tmp_dir/input1.txt
echo "$input2" > $tmp_dir/input2.txt


case="標準入力から受け取った内容を集計できること"

want=$(cat <<EOF
2	foo	[STDIN:1 STDIN:3]
2	baz	[STDIN:4 STDIN:5]
EOF
)

## test
got=$(go run dup.go < $tmp_dir/input1.txt)

## check
diff -u <(echo "$want" | sort) <(echo "$got" | sort)
if [ $? -eq 0 ]; then
  echo "$case: ok ✅"
else
  echo "$case: failed ❌"
fi


case="ファイルから入力した内容を集計できること"

want=$(cat <<EOF
3	foo	[$tmp_dir/input1.txt:1 $tmp_dir/input1.txt:3 $tmp_dir/input2.txt:1]
3	bar	[$tmp_dir/input1.txt:2 $tmp_dir/input2.txt:2 $tmp_dir/input2.txt:3]
3	baz	[$tmp_dir/input1.txt:4 $tmp_dir/input1.txt:5 $tmp_dir/input2.txt:4]
EOF
)

## test
got=$(go run dup.go $tmp_dir/input1.txt $tmp_dir/input2.txt)

## check
diff -u <(echo "$want" | sort) <(echo "$got" | sort)
if [ $? -eq 0 ]; then
  echo "$case: ok ✅"
else
  echo "$case: failed ❌"
fi
