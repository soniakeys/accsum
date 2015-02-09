// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

// Accsum contains routines for accurate floating point arithmetic.
//
// These are mostly implementations of algorithms developed by Siegfried M.
// Rump and colleagues.  See
// http://www.ti3.tu-harburg.de/rump/Research/topics.php#dot_product
// for an overview and links to published papers.  See also
// http://www.ti3.tu-harburg.de/rump/AccSum.zip for reference versions of most
// of these algorithms in Matlab.
//
// Very roughly, this graph shows some relationships of speed and accuracy.
//
//  Accuracy, or capability
//   ^
//   | NearSum
//   |      AccSumHuge
//   |            AccSum
//   |                 AccSumK
//   |                        PrecSum
//   |                             PriestSum
//   |                                  SumK, SumKVert
//   |                                           Sum2
//   |                                             XSum
//   |                                              KahanB
//   |                                               KahanSum
//   |                                                     PairSum
//   |
//   |                                                          Sum
//   +-------------------------------------------------------------->  Speed
//
// Well, I haven't benchmarked any of this yet.  Some of that may be wrong,
// but that's the basic idea.  PairSum is a big improvement over Sum if you
// have to sum a bunch of numbers.  Sum2 is the first Rump algorithm; it is
// fast and has has provable qualities.  AccSum gives great accuracy with
// reasonable speed.  NearSum gives the true round-to-nearest result, although
// at cost of time.
package accsum
