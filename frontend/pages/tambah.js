import React from 'react'
import Home from '.'
import HomeButton from '../components/HomeButton'
import TambahPegawai from '../components/TambahPegawai'


function tambah() {
  return (
    <div>
      <TambahPegawai></TambahPegawai>
      <HomeButton></HomeButton>
    </div>
  )
}

export default tambah
