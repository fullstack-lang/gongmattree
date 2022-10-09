// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { SliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield-db';

// insertion point for imports
import { GongStructDB } from './gongstruct-db'

@Injectable({
  providedIn: 'root'
})
export class SliceOfPointerToGongStructFieldService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  SliceOfPointerToGongStructFieldServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private sliceofpointertogongstructfieldsUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.sliceofpointertogongstructfieldsUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/sliceofpointertogongstructfields';
  }

  /** GET sliceofpointertogongstructfields from the server */
  getSliceOfPointerToGongStructFields(): Observable<SliceOfPointerToGongStructFieldDB[]> {
    return this.http.get<SliceOfPointerToGongStructFieldDB[]>(this.sliceofpointertogongstructfieldsUrl)
      .pipe(
        tap(_ => this.log('fetched sliceofpointertogongstructfields')),
        catchError(this.handleError<SliceOfPointerToGongStructFieldDB[]>('getSliceOfPointerToGongStructFields', []))
      );
  }

  /** GET sliceofpointertogongstructfield by id. Will 404 if id not found */
  getSliceOfPointerToGongStructField(id: number): Observable<SliceOfPointerToGongStructFieldDB> {
    const url = `${this.sliceofpointertogongstructfieldsUrl}/${id}`;
    return this.http.get<SliceOfPointerToGongStructFieldDB>(url).pipe(
      tap(_ => this.log(`fetched sliceofpointertogongstructfield id=${id}`)),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>(`getSliceOfPointerToGongStructField id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new sliceofpointertogongstructfield to the server */
  postSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB): Observable<SliceOfPointerToGongStructFieldDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    sliceofpointertogongstructfielddb.GongStruct = new GongStructDB
    let _GongStruct_SliceOfPointerToGongStructFields_reverse = sliceofpointertogongstructfielddb.GongStruct_SliceOfPointerToGongStructFields_reverse
    sliceofpointertogongstructfielddb.GongStruct_SliceOfPointerToGongStructFields_reverse = new GongStructDB

    return this.http.post<SliceOfPointerToGongStructFieldDB>(this.sliceofpointertogongstructfieldsUrl, sliceofpointertogongstructfielddb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        sliceofpointertogongstructfielddb.GongStruct_SliceOfPointerToGongStructFields_reverse = _GongStruct_SliceOfPointerToGongStructFields_reverse
        this.log(`posted sliceofpointertogongstructfielddb id=${sliceofpointertogongstructfielddb.ID}`)
      }),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>('postSliceOfPointerToGongStructField'))
    );
  }

  /** DELETE: delete the sliceofpointertogongstructfielddb from the server */
  deleteSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB | number): Observable<SliceOfPointerToGongStructFieldDB> {
    const id = typeof sliceofpointertogongstructfielddb === 'number' ? sliceofpointertogongstructfielddb : sliceofpointertogongstructfielddb.ID;
    const url = `${this.sliceofpointertogongstructfieldsUrl}/${id}`;

    return this.http.delete<SliceOfPointerToGongStructFieldDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted sliceofpointertogongstructfielddb id=${id}`)),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>('deleteSliceOfPointerToGongStructField'))
    );
  }

  /** PUT: update the sliceofpointertogongstructfielddb on the server */
  updateSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB): Observable<SliceOfPointerToGongStructFieldDB> {
    const id = typeof sliceofpointertogongstructfielddb === 'number' ? sliceofpointertogongstructfielddb : sliceofpointertogongstructfielddb.ID;
    const url = `${this.sliceofpointertogongstructfieldsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    sliceofpointertogongstructfielddb.GongStruct = new GongStructDB
    let _GongStruct_SliceOfPointerToGongStructFields_reverse = sliceofpointertogongstructfielddb.GongStruct_SliceOfPointerToGongStructFields_reverse
    sliceofpointertogongstructfielddb.GongStruct_SliceOfPointerToGongStructFields_reverse = new GongStructDB

    return this.http.put<SliceOfPointerToGongStructFieldDB>(url, sliceofpointertogongstructfielddb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        sliceofpointertogongstructfielddb.GongStruct_SliceOfPointerToGongStructFields_reverse = _GongStruct_SliceOfPointerToGongStructFields_reverse
        this.log(`updated sliceofpointertogongstructfielddb id=${sliceofpointertogongstructfielddb.ID}`)
      }),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>('updateSliceOfPointerToGongStructField'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
