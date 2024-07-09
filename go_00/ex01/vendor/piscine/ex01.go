/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex01.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: miyuu <miyuu@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/28 19:13:46 by miyuu             #+#    #+#             */
/*   Updated: 2024/06/29 22:06:38 by miyuu            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func AlphaDown() {
	for c := 'z'; c >= 'a'; c-- {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

// func AlphaDown() {
// 	var result string
// 	for c := 'a'; c <= 'z'; c++ {
// 		// ft.PrintRune("%c", c)
// 		result = string(c) + result
// 	}
// 	fmt.Println(result)
// }

//string = 文字列型 。string(c)でcを文字列に変換している。string(c) + resultでcをresultの前に追加している。
