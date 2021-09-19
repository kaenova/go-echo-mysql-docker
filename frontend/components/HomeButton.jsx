import React, { Component } from 'react'
import { AiOutlineHome } from "react-icons/ai";
import Link from "next/link"

export default class HomeButton extends Component {
  render() {
    return (
      <Link href="/">
        <a href="/">
          <span className="fixed z-10 bg-black p-1 rounded-md shadow-lg bottom-10 right-10">
            <AiOutlineHome size="3rem" color="white" />
          </span>
        </a>
      </Link>
    )
  }
}
