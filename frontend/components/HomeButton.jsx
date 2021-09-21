import React, { Component } from 'react'
import { AiOutlineHome } from "react-icons/ai";
import Link from "next/link"
import {motion} from 'framer-motion'

export default class HomeButton extends Component {
  render() {
    return (
      <Link href="/">
        <a href="/">
          <motion.span className="fixed z-10 bg-black p-1 rounded-md shadow-lg bottom-10 right-10"
          whileHover={{scale:1.2}}
          whileTap={{scale:0.9}}>
            <AiOutlineHome size="2rem" color="white" />
          </motion.span>
        </a>
      </Link>
    )
  }
}
