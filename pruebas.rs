let v: Vec<i64> =Vec::with_capacity(10);
let b: Vec<i64> =Vec::new();
println!("{}",v[0][0]);

let b:Vec<Vec<i64>> = vec![vec![1; 10],vec![2; 8],vec![3; 15],vec![5; 2],vec![8; 1]];
println!("{}",b[0]);
println!("{}",b[1]);
println!("{}",b[2]);
println!("{}",b[3]);
println!("{}",b[4]);

let b:Vec<Vec<i64>> = vec![vec![1; 10],vec![2; 8],vec![3; 15],vec![5; 2],vec![8; 1]];
println!("{}",b[0][0]);
let v:Vec<i64> = vec![1,2,3,4,5];
println!("{}",v[0]);
println!("{}",v[1]);
println!("{}",v[2]);
println!("{}",v[3]);
println!("{}",v[4]);

println!("{}",vec![vec![1; 10],vec![2; 8],vec![3; 15],vec![5; 2],vec![8; 1]]);

let mut arr3: [[[i64; 4];2]; 2] = [
[ [ 1, 3, 5, 7], [ 9, 11, 13, 15] ],
[ [ 2, 4, 6, 8], [10, 12, 14, 16] ]
];
println!("{}",arr3);
arr3[0][0][0]=10;
println!("{}",arr3);

fn main(){
    println!("----------------------");
    println!("----ARCHIVO BASICO----");
    println!("----------------------");
    let mut bo1: bool = false;
    let mut a1:&str="hola";
    let mut a2:char='c';
    let mut a3:i64=4;
    let  mut a4:f64=5.0;
    let mut a5:usize=7;
    println!("{}",bo1);
    println!("{}",a1);
    println!("{}",a2);
    println!("{}",a3);
    println!("{}",a4);
    println!("{}",a5);
    bo1 = true;
    a1="adios";
    a2='b';
    a3=5;
    a4=56.0;
//    let a5:usize=7;
    println!("{}",bo1);
    println!("{}",a1);
    println!("{}",a2);
    println!("{}",a3);
    println!("{}",a4);
    println!("{}",a5);
}