use std::collections::HashSet;
use std::fs;

use itertools::Itertools;

type Point = [i32; 4];

struct Grid {
    cubes: HashSet<Point>,
    directions: [Point; 80],
}

impl Grid {
    fn simulate(&mut self, iterations: i32) {
        let mut cubes_copy: HashSet<Point> = HashSet::new();
        let mut potentially_activated: HashSet<Point> = HashSet::new();

        for _ in 0..iterations {
            cubes_copy.clone_from(&self.cubes);
            potentially_activated.clear();

            for c in self.cubes.iter() {
                let (num_active, inactive) = self.neighbors(c, true);
                if num_active != 2 && num_active != 3 {
                    cubes_copy.remove(c);
                }
                potentially_activated.extend(&inactive);
            }

            for c in potentially_activated.iter() {
                let (num_active, _) = self.neighbors(c, false);
                if num_active == 3 {
                    cubes_copy.insert(*c);
                }
            }

            self.cubes.clone_from(&cubes_copy);
        }
    }

    fn neighbors(&self, c: &Point, get_inactive: bool) -> (i32, HashSet<Point>) {
        let mut num_active: i32 = 0;
        let mut inactive = HashSet::new();

        for d in self.directions.iter() {
            let p: Point = [c[0] + d[0], c[1] + d[1], c[2] + d[2], c[3] + d[3]];
            if self.cubes.contains(&p) {
                num_active += 1;
            } else if get_inactive {
                inactive.insert(p);
            }
        }

        (num_active, inactive)
    }

    fn new(s: String) -> Grid {
        let mut cubes: HashSet<Point> = HashSet::new();

        let (mut i, mut j) = (0, 0);
        for c in s.chars() {
            match c {
                '#' => {
                    cubes.insert([i, j, 0, 0]);
                    j += 1
                }
                '.' => j += 1,
                '\n' => {
                    j = 0;
                    i += 1
                }
                _ => {
                    println!("Invalid character: {}", c)
                }
            }
        }

        let mut directions: [Point; 80] = [[0, 0, 0, 0]; 80];
        let mut cnt = 0;
        for i in (-1..2)
            .cartesian_product(-1..2)
            .cartesian_product(-1..2)
            .cartesian_product(-1..2)
        {
            let p: Point = [i.0 .0 .0, i.0 .0 .1, i.0 .1, i.1];
            if p[0] != 0 || p[1] != 0 || p[2] != 0 || p[3] != 0 {
                directions[cnt] = p;
                cnt += 1
            }
        }

        Grid { cubes, directions }
    }
}

fn main() {
    let t = fs::read_to_string("input.txt").expect("Unable to read file");
    let mut g = Grid::new(t);
    g.simulate(6);
    println!("{}", g.cubes.len())
}
