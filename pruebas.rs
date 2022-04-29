let v: Vec<i64> =Vec::with_capacity(10);
let b: Vec<i64> =Vec::new();
println!("{}",v[0][0]);

let v:Vec<Vec<i64>> = vec![vec![1; 10],vec![2; 8],vec![3; 15],vec![5; 2],vec![8; 1]];
println!("{}",v[0][0]);

let b:Vec<Vec<i64>> = vec![vec![1; 10],vec![2; 8],vec![3; 15],vec![5; 2],vec![8; 1]];
println!("{}",b[0][0]);
let v:Vec<i64> = vec![1,2,3,4,5];
println!("{}",v[0]);
println!("{}",v[1]);
println!("{}",v[2]);
println!("{}",v[3]);
println!("{}",v[4]);

println!("{}",vec![vec![1; 10],vec![2; 8],vec![3; 15],vec![5; 2],vec![8; 1]]);