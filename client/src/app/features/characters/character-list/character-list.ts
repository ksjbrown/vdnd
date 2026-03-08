import { Component, OnInit } from '@angular/core';
import { Character, CharactersService } from '../characters';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-character-list',
  imports: [],
  templateUrl: './character-list.html',
  styleUrl: './character-list.css',
})
export class CharacterList implements OnInit {

  characters: Character[] = []

  constructor(private charactersService: CharactersService) { }

  ngOnInit(): void {
    this.charactersService.getCharacters().subscribe((c) => this.characters = c)
  }

}
